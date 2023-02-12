package ordershandlers

import (
	"encoding/json"
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
		http.Error(w, errors.New("wrong content-type").Error(), http.StatusBadRequest)
		log.Print("post.order:wrong content-type")
		return
	}

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("post.order: %v", err)
		return
	}

	orderNumber := string(body)

	if orderNumber == "" {
		http.Error(w, errors.New("empty body").Error(), http.StatusBadRequest)
		log.Print("empty body")
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

		log.Printf("post.order.service %v", err)

		if errors.Is(err, server.ErrUploadedByOtherUser) {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}

		if errors.Is(err, server.ErrUploadedByUser) {
			http.Error(w, err.Error(), http.StatusOK)
			return
		}

		if errors.Is(err, server.ErrOrderNumberFormat) {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
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
		http.Error(w, errors.New("Unauthorized").Error(), http.StatusUnauthorized)
		return
	}

	orders, err := gh.Service.GetOrdersInformation(ctx, userID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(orders) == 0 {
		http.Error(w, "нет данных для ответа", http.StatusNoContent)
		return
	}

	body, err := json.Marshal(orders)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
