package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rautaruukkipalich/go_auth/internal/model"
	"github.com/rautaruukkipalich/go_auth/internal/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestServer_HandleRegister(t *testing.T) { //nolint: no cover

	s := newServer(teststore.New(), "info")

	testCases := []struct {
		name string
		payload interface{}
		expectedCode int
	}{
		{
			name: "valid registration",
			payload: map[string]string{
				"username": "adsadsadasda",
				"password": "321",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "bad request",
			payload: "123",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "bad username",
			payload: map[string]string{
				"username": "11111",
				"password": "321",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "bad password",
			payload: map[string]string{
				"username": "adsadsadasda",
				"password": " ",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rr := httptest.NewRecorder()

			var json_data bytes.Buffer
			json.NewEncoder(&json_data).Encode(tc.payload)

			req, _ := http.NewRequest(http.MethodPost, "/register", &json_data)
			s.ServeHTTP(rr, req)

			assert.Equal(t, tc.expectedCode, rr.Code)
		})
	}
}


func TestServer_HandleAuth(t *testing.T) { //nolint: errcheck

	s := newServer(teststore.New(), "info")
	u := model.TestUser(t)
	s.store.User().Create(u)

	testCases := []struct {
		name string
		payload interface{}
		expectedCode int
	}{
		{
			name: "valid auth",
			payload: map[string]string{
				"username": u.Username,
				"password": u.Password,
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "invalid payload",
			payload: "123",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid username",
			payload: map[string]string{
				"username": "123",
				"password": u.Password,
			},
			expectedCode: http.StatusUnauthorized, 
		},
		{
			name: "invalid password",
			payload: map[string]string{
				"username": u.Username,
				"password": "132",
			},
			expectedCode: http.StatusUnauthorized,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			rr := httptest.NewRecorder()

			var json_data bytes.Buffer
			json.NewEncoder(&json_data).Encode(tc.payload)

			req, _ := http.NewRequest(http.MethodPost, "/auth", &json_data)
			s.ServeHTTP(rr, req)

			assert.Equal(t, tc.expectedCode, rr.Code)
		})
	}
}