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

	response := models.RegistrationResponse{Result: true}

	resBody, err := json.Marshal(&response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
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

	response := models.LoginResponse{Token: "12345678"}

	resBody, err := json.Marshal(&response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resBody)
}
