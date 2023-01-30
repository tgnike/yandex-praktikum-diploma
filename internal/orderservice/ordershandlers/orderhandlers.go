package ordershandlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/tgnike/yandex-praktikum-diploma/internal/server"
)

type PostOrderHandler struct {
	*chi.Mux
	Service server.Orders
}

func (ph *PostOrderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("it works! PostOrderHandler"))
}

type GetOrdersHandler struct {
	*chi.Mux
	Service server.Orders
}

func (gh *GetOrdersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("it works! GetOrdersHandler"))
}
