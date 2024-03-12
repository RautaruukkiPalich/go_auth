package store

import "github.com/rautaruukkipalich/go_auth/internal/model"

type UserRepository interface {
	Create(*model.User) (*model.User, error)
	Auth(*model.User) (string, error)

	FindById(int) (*model.User, error)
	FindByUsername(string) (*model.User, error)
	FindBySlug(string) (*model.User, error)
	
	UpdatePassword(*model.User, string) (error)
	UpdateUsername(*model.User, string) (error) 
}