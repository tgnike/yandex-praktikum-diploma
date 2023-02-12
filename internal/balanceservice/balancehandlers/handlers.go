package balancehandlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/tgnike/yandex-praktikum-diploma/internal/models"
	"github.com/tgnike/yandex-praktikum-diploma/internal/server"
)

// GET /api/user/balance — получение текущего баланса счёта баллов лояльности пользователя;
// POST /api/user/balance/withdraw — запрос на списание баллов с накопительного счёта в счёт оплаты нового заказа;
// GET /api/user/withdrawals

type GetBalanceHandler struct {
	*chi.Mux
	Service server.Balance
}

func (gbh *GetBalanceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	userID, ok := ctx.Value(server.UserContext).(models.UserID)
	if !ok {
		http.Error(w, errors.New("Unauthorized").Error(), http.StatusUnauthorized)
		return
	}

	balance, err := gbh.Service.GetBalance(ctx, &userID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := json.Marshal(balance)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)

}

type WithdrawRequestHandler struct {
	*chi.Mux
	Service server.Balance
}

func (wt *WithdrawRequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	userID, ok := ctx.Value(server.UserContext).(models.UserID)
	if !ok {
		http.Error(w, errors.New("Unauthorized").Error(), http.StatusUnauthorized)
		return
	}

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	withdrawal := &models.WithdrawalRequest{}

	if err := json.Unmarshal(body, &withdrawal); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = wt.Service.WithdrawRequest(ctx, withdrawal, &userID)

	if err != nil {

		if errors.Is(err, server.ErrOrderNumberFormat) {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(""))
}

type GetWithdrawalsHandler struct {
	*chi.Mux
	Service server.Balance
}

func (gwh *GetWithdrawalsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	userID, ok := ctx.Value(server.UserContext).(models.UserID)
	if !ok {
		http.Error(w, errors.New("Unauthorized").Error(), http.StatusUnauthorized)
		return
	}

	withdrawals, err := gwh.Service.Withdrawals(ctx, &userID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := json.Marshal(withdrawals)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
