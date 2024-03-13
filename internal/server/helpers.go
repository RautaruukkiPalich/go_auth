package server

import (
	"encoding/json"
	"net/http"

	"github.com/rautaruukkipalich/go_auth/internal/model"
)

func (s *Server) getFormFromBody(r *http.Request, form interface{}) (err error) {
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		return err
	}
	return nil
}

func (s *Server) getUserFromContext(r *http.Request) (*model.User, error) {
	ctx := r.Context()
	user, err := s.store.User().GetById(ctx.Value(UserKey).(int))

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Server) error(w http.ResponseWriter, r *http.Request, err errorResponse) {
	s.logger.Printf("error: %v", err)
	s.respond(w, r, err.Code, err)
}

func (s *Server) respond(w http.ResponseWriter, r *http.Request, code int, data any) {
	w.WriteHeader(code)
	if data != nil {
		w.Header().Add("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			s.logger.Printf("error: %v", err)
		}
	}
}
