package server

import (
	"database/sql"
	"net/http"
	"github.com/rs/cors"

	"github.com/rautaruukkipalich/go_auth/internal/store/sqlstore"
)

func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	
	if err != nil {
		return err
	}

	defer db.Close()

	store := sqlstore.New(db)
	srv := newServer(store)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{
			"http://localhost:8080",
			"http://127.0.0.1:8080",
		},
		AllowedMethods:   []string{
			http.MethodGet, 
			http.MethodPost, 
			http.MethodPatch,
		},
		AllowCredentials: true,
	})

	handler := c.Handler(srv.router)

	return http.ListenAndServe(config.BindAddress, handler)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	
	return db, nil
}
