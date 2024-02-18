package middleware

import (
	"context"
	"log"
	"net/http"
	"github.com/rautaruukkipalich/go_auth/internal/model"
)

type userKey string
const UserKey userKey = "user"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func (w http.ResponseWriter, r *http.Request) {
			user, err := checkUser(r)
			if err != nil {
				log.Println(err)
			}
			if user.Id == 0 {
				http.Error(
					w,
					http.StatusText(http.StatusUnauthorized),
					http.StatusUnauthorized,
				)
			}
			next.ServeHTTP(
				w, r.WithContext(context.WithValue(r.Context(), user, UserKey)),
			)
		},
	)
}

func checkUser(r *http.Request) (*model.User, error) {
	return nil, nil
}