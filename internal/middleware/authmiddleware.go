package middleware

import (
	"context"
	"log"
	"net/http"
)

type userKey string
const UserKey userKey = "userId"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func (w http.ResponseWriter, r *http.Request) {
			userId, err := parseToken(r)
			if err != nil {
				log.Println(err)
			}
			if userId == 0 {
				http.Error(
					w,
					http.StatusText(http.StatusUnauthorized),
					http.StatusUnauthorized,
				)
				return
			}
			next.ServeHTTP(
				w, r.WithContext(context.WithValue(r.Context(), userId, UserKey)),
			)
		},
	)
}

func parseToken(r *http.Request) (int, error) {
	// check user header fields
	// decode header
	// check expire
	// return UserId by token credentials
	return 0, nil
}