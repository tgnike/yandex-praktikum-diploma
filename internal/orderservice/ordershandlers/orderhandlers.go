package ordershandlers

import (
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/tgnike/yandex-praktikum-diploma/internal/models"
	"github.com/tgnike/yandex-praktikum-diploma/internal/server"
)

type PostOrderHandler struct {
	*chi.Mux
	Service server.Orders
	UserID  models.UserID
}

func (ph *PostOrderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	ct := r.Header.Get("Content-type")

	if ct != "text/plain" {
		http.Error(w, errors.New("Wrong content-type").Error(), http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	orderNumber := string(body)

	if orderNumber == "" {
		http.Error(w, errors.New("empty body").Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	userID, ok := ctx.Value(server.UserContext).(models.UserID)
	if !ok {
		http.Error(w, errors.New("Unauthorized").Error(), http.StatusUnauthorized)
		return
	}

	err = ph.Service.PostOrder(ctx, orderNumber, userID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(""))
}

type GetOrdersHandler struct {
	*chi.Mux
	Service server.Orders
	UserID  models.UserID
}

func (gh *GetOrdersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID, ok := ctx.Value(server.UserContext).(models.UserID)
	if !ok {
		http.Error(w, errors.New("empty body").Error(), http.StatusUnauthorized)
		return
	}

	log.Print(userID)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("it works! GetOrdersHandler"))
}
