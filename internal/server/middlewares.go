package server

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/rautaruukkipalich/go_auth/internal/utils/jwt"
)

var ErrInvalidToken = errors.New("invalid token")

func (s *Server) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			authHeaderChunks := strings.Split(authHeader, " ")

			if len(authHeaderChunks) != 2 {
				s.error(w, r, errorResponse{Error: ErrInvalidToken.Error(), Code: http.StatusUnauthorized})
				return
			}

			userId, err := jwt.DecodeJWTToken(authHeaderChunks[1])

			if err != nil || userId == 0 {
				s.error(w, r, errorResponse{Error: ErrInvalidToken.Error(), Code: http.StatusUnauthorized})
				return
			}

			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), UserKey, userId)))
		},
	)
}