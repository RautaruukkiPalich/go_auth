package store

import "github.com/rautaruukkipalich/go_auth/internal/model"

type IUserRepository interface {
	Create(*model.User) (*model.User, error)
	Auth(*model.User) (string, error)
	FindById(int) (*model.User, error)
	Find(string) (*model.User, error)
}