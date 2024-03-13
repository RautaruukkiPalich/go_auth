package server

import (
	"net/http"
	"time"

	"github.com/rautaruukkipalich/go_auth/internal/model"
)

type (
	userKey string

	userAuthRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	userResponse struct {
		ID                 int       `json:"id"`
		Username           string    `json:"username"`
		CreatedAt          time.Time `json:"created_at"`
		LastPasswordChange time.Time `json:"last_password_change"`
	}

	loginResponse struct {
		AccessToken string `json:"token"`
	}

	changeUsername struct {
		Username string `json:"username"`
	}

	changePassword struct {
		Password string `json:"password"`
	}

	errorResponse struct {
		Code  int    `json:"code"`
		Error string `json:"error"`
	}
)

const (
	UserKey userKey = "userID"
)

func responseFromUser(user *model.User) *userResponse {
	return &userResponse{
		ID:                 user.Id,
		Username:           user.Username,
		CreatedAt:          user.CreatedAt,
		LastPasswordChange: user.LastPasswordChange,
	}
}

// @Summary		Register
// @Description	register by username and password
// @Tags			auth
// @Accept			json
// @Produce		json
// @Param			input	body		userAuthRequest	true	"register"
// @Failure		400,404	{object}	errorResponse
// @Success		500		{object}	errorResponse
// @Success		default	{object}	errorResponse
// @Router			/register [post]
func (s *Server) Register() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		request := &userAuthRequest{}

		if err := s.getFormFromBody(r, request); err != nil {
			s.error(w, r, errorResponse{Error: err.Error(), Code: http.StatusBadRequest})
			return
		}

		u := &model.User{
			Username: request.Username,
			Password: request.Password,
		}

		_, err := s.store.User().Create(u)
		if err != nil {
			s.error(w, r, errorResponse{Error: err.Error(), Code: http.StatusUnprocessableEntity})
			return
		}

		s.respond(w, r, http.StatusCreated, nil)
	}
}

// @Summary		Login
// @Description	get token by username and password
// @Tags			auth
// @Accept			json
// @Produce		json
// @Param			input	body		userAuthRequest	true	"login"
// @Success		200		{object}	loginResponse
// @Failure		400,404	{object}	errorResponse
// @Success		500		{object}	errorResponse
// @Success		default	{object}	errorResponse
// @Router			/auth [post]
func (s *Server) Auth() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var form userAuthRequest

		if err := s.getFormFromBody(r, &form); err != nil {
			s.error(w, r, errorResponse{Error: err.Error(), Code: http.StatusBadRequest})
			return
		}

		u := &model.User{
			Username: form.Username,
			Password: form.Password,
		}

		token, err := s.store.User().Auth(u)

		if err != nil {
			s.error(w, r, errorResponse{Error: err.Error(), Code: http.StatusUnauthorized})
			return
		}

		s.respond(w, r, http.StatusOK, loginResponse{AccessToken: token})
	}
}

// @Summary		Account
// @Security		ApiKeyAuth
// @Description	account info
// @Tags			me
// @Accept			json
// @Produce		json
// @Success		200		{object}	userResponse
// @Failure		400,404	{object}	errorResponse
// @Success		500		{object}	errorResponse
// @Success		default	{object}	errorResponse
// @Router			/me [get]
func (s *Server) Me() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		user, err := s.getUserFromContext(r)
		if err != nil {
			s.error(w, r, errorResponse{Error: err.Error(), Code: http.StatusBadRequest})
			return
		}

		s.respond(w, r, http.StatusOK, responseFromUser(user))
	}
}

// @Summary		Edit Password
// @Security		ApiKeyAuth
// @Description	change password
// @Tags			me
// @Accept			json
// @Produce		json
// @Param			input	body		changePassword	true	"editPassword"
// @Failure		400,404	{object}	errorResponse
// @Success		500		{object}	errorResponse
// @Success		default	{object}	errorResponse
// @Router			/me/password [patch]
func (s *Server) EditPassword() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var form changePassword

		if err := s.getFormFromBody(r, &form); err != nil {
			s.error(w, r, errorResponse{Error: err.Error(), Code: http.StatusUnprocessableEntity})
			return
		}

		user, err := s.getUserFromContext(r)
		if err != nil {
			s.error(w, r, errorResponse{Error: err.Error(), Code: http.StatusBadRequest})
			return
		}

		if err = s.store.User().SetPassword(user, form.Password); err != nil {
			s.error(w, r, errorResponse{Error: err.Error(), Code: http.StatusUnprocessableEntity})
			return
		}

		s.respond(w, r, http.StatusOK, nil)
	}
}

// @Summary		Edit Username
// @Security		ApiKeyAuth
// @Description	change username
// @Tags			me
// @Accept			json
// @Produce		json
// @Param			input	body		changeUsername	true	"editusername"
// @Failure		400,404	{object}	errorResponse
// @Success		500		{object}	errorResponse
// @Success		default	{object}	errorResponse
// @Router			/me [patch]
func (s *Server) EditUsername() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var form changeUsername

		if err := s.getFormFromBody(r, form); err != nil {
			s.error(w, r, errorResponse{Error: err.Error(), Code: http.StatusUnprocessableEntity})
			return
		}

		user, err := s.getUserFromContext(r)
		if err != nil {
			s.error(w, r, errorResponse{Error: err.Error(), Code: http.StatusBadRequest})
			return
		}

		if err = s.store.User().SetUsername(user, form.Username); err != nil {
			s.error(w, r, errorResponse{Error: err.Error(), Code: http.StatusUnprocessableEntity})
			return
		}

		s.respond(w, r, http.StatusOK, nil)
	}
}
