package store

import "github.com/rautaruukkipalich/go_auth/internal/model"

type UserRepositorier interface {
	UserCreater
	UserAuthenticater
	UserGetter
	UserSetter
}

type UserCreater interface {
	Create(*model.User) (*model.User, error)
}

type UserAuthenticater interface {
	Auth(*model.User) (string, error)
}

type UserGetter interface {
	GetById(int) (*model.User, error)
	GetByUsername(string) (*model.User, error)
	GetBySlug(string) (*model.User, error)
}

type UserSetter interface {
	SetPassword(*model.User, string) (error)
	SetUsername(*model.User, string) (error) 
}