package store

type Storer interface {
	User() UserRepositorier
}
