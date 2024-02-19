package store

type Store interface {
	User() IUserRepository
}