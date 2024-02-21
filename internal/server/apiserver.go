package server

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rautaruukkipalich/go_auth/internal/store/sqlstore"
)

func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()

	srv := newServer(
		sqlstore.New(db),
		config.LogLevel,
	)

	if err := migrateTables(db); err != nil {
		srv.logger.Error(err)
	}

	c := newCors(
		NewCorsConfig(),
	)

	server := &http.Server{
		Addr: config.BindAddress,
		Handler: c.Handler(srv.router),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	srv.logger.Info(fmt.Sprintf("server up on '%s'; log level '%s'",
		server.Addr,
		srv.logger.Level,
		),
	)

	return server.ListenAndServe()
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

func migrateTables(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", 
		driver,
	)
	if err != nil {
		fmt.Println("err :", err)
		return err
	}

	m.Up()

	return nil
}
