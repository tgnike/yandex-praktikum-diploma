package authhandlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/tgnike/yandex-praktikum-diploma/internal/models"
	"github.com/tgnike/yandex-praktikum-diploma/internal/server"
)

type RegistationHandler struct {
	*chi.Mux
	Service server.Users
}

func (rh *RegistationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var userJSON models.UserJSON

	if err := json.Unmarshal(body, &userJSON); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	token, err := rh.Service.Register(ctx, &userJSON)

	if err != nil {
		// TODO типы ошибок
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	response := models.RegistrationResponse{Result: true}

	resBody, err := json.Marshal(&response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Authorization", "Token "+string(token))
	w.WriteHeader(http.StatusOK)
	w.Write(resBody)
}

type LoginHandler struct {
	*chi.Mux
	Service server.Users
}

func (lh *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var userJSON models.UserJSON

	if err := json.Unmarshal(body, &userJSON); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	token, err := lh.Service.Login(ctx, &userJSON)

	if err != nil {
		// TODO типы ошибок
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	response := models.LoginResponse{Token: string(token)}

	resBody, err := json.Marshal(&response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Authorization", "Token "+string(token))
	w.WriteHeader(http.StatusOK)
	w.Write(resBody)
}
