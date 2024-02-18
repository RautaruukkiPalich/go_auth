package server

import (
	"fmt"
	"net/http"

	"github.com/rautaruukkipalich/go_auth/internal/store"
	"github.com/rautaruukkipalich/go_auth/internal/transport/rest"
	mw "github.com/rautaruukkipalich/go_auth/internal/middleware"
	"github.com/sirupsen/logrus"
)

type APIserver struct {
	config *Config
	logger *logrus.Logger
	router *http.ServeMux
	store  *store.Store
}

func New(config *Config) *APIserver {
	return &APIserver{
		config: config,
		logger: logrus.New(),
		router: http.NewServeMux(),
	}
}

func (s *APIserver) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	if err := s.configureStore(); err != nil {
		return err
	}

	s.configureRouter()

	s.logger.Info(
		fmt.Sprintf("Starting API server at %s", s.config.BindAddress),
	)
	return http.ListenAndServe(s.config.BindAddress, s.router)
}

func (s *APIserver) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)
	return nil
}


func (s *APIserver) configureStore() error {
	st := store.New(s.config.Store)
	err := st.Open()
	if err != nil {
		return err
	}
	s.store = st
	return nil
}

func (s *APIserver) configureRouter() {
	s.router.HandleFunc(
		"/register", 
		rest.Register,
	)
	s.router.HandleFunc(
		"/auth", 
		rest.Auth,
	)
	s.router.Handle(
		"/me",
		mw.AuthMiddleware(
			http.HandlerFunc(rest.Me),
		),
	)
	s.router.Handle(
		"/me/pass", 
		mw.AuthMiddleware(
			http.HandlerFunc(rest.EditPassword),
		),
	)
	s.router.Handle(
		"/me/username", 
		mw.AuthMiddleware(
			http.HandlerFunc(rest.EditUsername),
		),
	)
}
