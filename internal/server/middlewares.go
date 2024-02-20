package server

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/rautaruukkipalich/go_auth/internal/utils"
)

func (s *Server) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			authHeaderChunks := strings.Split(authHeader, " ")
			err_msg := "invalid token"

			if len(authHeaderChunks) != 2 {
				s.logger.Error(err_msg)
				s.error(w, r, http.StatusUnauthorized, errors.New(err_msg))
				return
			}

			userId, err := utils.DecodeJWTToken(authHeaderChunks[1])

			if err != nil || userId == 0 {
				s.logger.Error(err_msg)
				s.error(w, r, http.StatusUnauthorized, errors.New(err_msg))
				return
			}

			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), UserKey, userId)))
		},
	)
}