package sqlstore

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/rautaruukkipalich/go_auth/internal/store"
)

type Store struct {
	db             *sql.DB
	userRepository *UserRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) User() store.IUserRepository {
	if s.userRepository == nil {
		s.userRepository = &UserRepository{
			sqlstore: s,
		}
	}
	return s.userRepository
}
