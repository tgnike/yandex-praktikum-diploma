package authhandlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	}

	tests := []struct {
		name    string
		handler *RegistationHandler
		args    args
		want    want
	}{
		{
			name:    "postitve test #1 200",
			handler: &RegistationHandler{},
			args:    args{r: registrationRequest, w: httptest.NewRecorder()},
			want:    want{StatusCode: http.StatusOK, Body: `{"result": true}`, ContentType: "application/json"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.handler.ServeHTTP(tt.args.w, tt.args.r)

			assert.Equal(t, tt.want.StatusCode, tt.args.w.Code)
			assert.Equal(t, tt.want.ContentType, tt.args.w.Header().Get("Content-type"))
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
	}

	tests := []struct {
		name    string
		handler *LoginHandler
		args    args
		want    want
	}{
		{
			name:    "postitve test #1 200",
			handler: &LoginHandler{},
			args:    args{r: registrationRequest, w: httptest.NewRecorder()},
			want:    want{StatusCode: http.StatusOK, Body: `{"token": "12345678"}`, ContentType: "application/json"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.handler.ServeHTTP(tt.args.w, tt.args.r)

			assert.Equal(t, tt.want.StatusCode, tt.args.w.Code)
			assert.Equal(t, tt.want.ContentType, tt.args.w.Header().Get("Content-type"))
			assert.JSONEq(t, tt.want.Body, tt.args.w.Body.String())

		})
	}
}
