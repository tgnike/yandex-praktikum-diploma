package authhandlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/tgnike/yandex-praktikum-diploma/internal/server"
)

type RegistationHandler struct {
	*chi.Mux
	Service server.Users
}

func (rh *RegistationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("it works! RegistationHandler"))
}

type LoginHandler struct {
	*chi.Mux
	Service server.Users
}

func (lh *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("it works! LoginHandler"))
}
