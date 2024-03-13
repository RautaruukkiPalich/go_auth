package teststore

import (
	"github.com/rautaruukkipalich/go_auth/internal/model"
	"github.com/rautaruukkipalich/go_auth/internal/store"
)

type Store struct {
	userRepository *UserRepository
}

func New() *Store {
	store := &Store{}
	store.userRepository = NewUserRepo(store)
	return store
}

func NewUserRepo(s *Store) *UserRepository {
	return &UserRepository{
		store: s,
		users: make(map[string]*model.User),
	}
}

func (s *Store) User() store.UserRepositorier {
	return s.userRepository
}
