package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rautaruukkipalich/go_auth/internal/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestServer_HandleIndex(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	s := newServer(teststore.New())
	s.ServeHTTP(rec, req)
	assert.Equal(t, rec.Code, http.StatusOK)
}
