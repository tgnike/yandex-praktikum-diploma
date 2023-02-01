package authhandlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tgnike/yandex-praktikum-diploma/internal/models"
)

type AuthTester struct {
}

func (at *AuthTester) Register(userJSON *models.UserJSON) (models.Token, error) {
	return models.Token("12345678"), nil
}
func (at *AuthTester) Login(userJSON *models.UserJSON) (models.Token, error) {
	return models.Token("12345678"), nil
}
func (at *AuthTester) CheckAuthToken(token string) (models.UserID, error) {
	return models.UserID("12345678"), nil
}

func TestRegistationHandler_ServeHTTP(t *testing.T) {

	jsonBody := []byte(`{"login": "hello", "password": "no"}`)
	bodyReader := bytes.NewReader(jsonBody)

	registrationRequest, err := http.NewRequest(http.MethodPost, "/api/user/register", bodyReader)

	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}

	type want struct {
		StatusCode  int
		Body        string
		ContentType string
		AuthHeader  string
	}

	tests := []struct {
		name    string
		handler *RegistationHandler
		args    args
		want    want
	}{
		{
			name:    "postitve test #1 200",
			handler: &RegistationHandler{Service: &AuthTester{}},
			args:    args{r: registrationRequest, w: httptest.NewRecorder()},
			want:    want{StatusCode: http.StatusOK, Body: `{"result": true}`, ContentType: "application/json", AuthHeader: "Token 12345678"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.handler.ServeHTTP(tt.args.w, tt.args.r)

			assert.Equal(t, tt.want.StatusCode, tt.args.w.Code)
			assert.Equal(t, tt.want.ContentType, tt.args.w.Header().Get("Content-type"))
			assert.Equal(t, tt.want.AuthHeader, tt.args.w.Header().Get("Authorization"))
			assert.JSONEq(t, tt.want.Body, tt.args.w.Body.String())

		})
	}
}

func TestLoginHandler_ServeHTTP(t *testing.T) {

	jsonBody := []byte(`{"login": "hello", "password": "no"}`)
	bodyReader := bytes.NewReader(jsonBody)

	registrationRequest, err := http.NewRequest(http.MethodPost, "/api/user/login", bodyReader)

	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}

	type want struct {
		StatusCode  int
		Body        string
		ContentType string
		AuthHeader  string
	}

	tests := []struct {
		name    string
		handler *LoginHandler
		args    args
		want    want
	}{
		{
			name:    "postitve test #1 200",
			handler: &LoginHandler{Service: &AuthTester{}},
			args:    args{r: registrationRequest, w: httptest.NewRecorder()},
			want:    want{StatusCode: http.StatusOK, Body: `{"token": "12345678"}`, ContentType: "application/json", AuthHeader: "Token 12345678"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.handler.ServeHTTP(tt.args.w, tt.args.r)

			assert.Equal(t, tt.want.StatusCode, tt.args.w.Code)
			assert.Equal(t, tt.want.ContentType, tt.args.w.Header().Get("Content-type"))
			assert.Equal(t, tt.want.AuthHeader, tt.args.w.Header().Get("Authorization"))
			assert.JSONEq(t, tt.want.Body, tt.args.w.Body.String())

		})
	}
}
