package server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	mw "github.com/rautaruukkipalich/go_auth/internal/middleware"
	"github.com/rautaruukkipalich/go_auth/internal/model"
	"github.com/rautaruukkipalich/go_auth/internal/store"
	"github.com/sirupsen/logrus"
)

type Server struct {
	router *mux.Router
	logger *logrus.Logger
	store  store.Store
}

func newServer(store store.Store) *Server {
	s := &Server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store:  store,
	}

	s.configureRouter()
	s.logger.Info("server up")
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) configureRouter() {
	s.router.HandleFunc("/register", s.Register()).Methods("POST")

	s.router.HandleFunc("/auth", s.Auth()).Methods("POST")

	s.router.Handle("/me", mw.AuthMiddleware(
		http.HandlerFunc(s.Me())),
	).Methods("GET")

	s.router.Handle("/me/pass", mw.AuthMiddleware(
		http.HandlerFunc(s.EditPassword())),
	).Methods("PATCH")

	s.router.Handle("/me/username", mw.AuthMiddleware(
		http.HandlerFunc(s.EditUsername())),
	).Methods("PATCH")
}

type userAuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type responseUser struct {
	Id                 int       `json:"id"`
	Username           string    `json:"username"`
	HashedPassword     string    `json:"hashed_password"`
	LastPasswordChange time.Time `json:"last_password_change"`
}

func responseFromUser(user *model.User) *responseUser {
	return &responseUser{
		Id:                 user.Id,
		Username:           user.Username,
		HashedPassword:     user.HashedPassword,
		LastPasswordChange: user.LastPasswordChange,
	}
}

type LoginResponse struct {
    AccessToken string `json:"access_token"`
}

func (s *Server) Register() http.HandlerFunc {
	// read data from request
	// check unique username
	// check password
	// hash password
	// save to database

	return func(w http.ResponseWriter, r *http.Request) {
		request := &userAuthRequest{}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u := &model.User{
			Username: request.Username,
			Password: request.Password,
		}

		user, err := s.store.User().Create(u)
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		resUser:= responseFromUser(user)
		
		s.respond(w, r, http.StatusCreated, resUser)
	}
}

func (s *Server) Auth() http.HandlerFunc {
	// read data from request
	// get user from DB by username
	// check password
	// return token
	return func(w http.ResponseWriter, r *http.Request) {
		request := &userAuthRequest{}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u := &model.User{
			Username: request.Username,
			Password: request.Password,
		}

		token, err := s.store.User().Auth(u)

		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		data := map[string]string{
			"token-type": "Bearer",
			"token": token,
		}

		s.respond(w, r, http.StatusCreated, data)

	}
}

func (s *Server) Me() http.HandlerFunc {
	// get user from DB
	// return user
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *Server) EditPassword() http.HandlerFunc {
	// read data from request
	// get user from DB
	// check password
	// hash password
	// save to database
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *Server) EditUsername() http.HandlerFunc {
	//read data from request
	// check unique username
	// save to database
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *Server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *Server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}

}
