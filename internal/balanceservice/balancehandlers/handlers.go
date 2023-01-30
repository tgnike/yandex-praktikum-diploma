package balancehandlers

import (
	"net/http"

	"github.com/go-chi/chi"
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

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("it works!"))

}

type WithdrawRequestHandler struct {
	*chi.Mux
	Service server.Balance
}

func (wt *WithdrawRequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("it works! WithdrawRequestHandler"))
}

type GetWithdrawalsHandler struct {
	*chi.Mux
	Service server.Balance
}

func (gwh *GetWithdrawalsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("it works! GetWithdrawalsHandler"))
}
