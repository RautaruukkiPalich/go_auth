package store

import (
	model "github.com/rautaruukkipalich/go_auth/internal/model"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	if err := r.store.db.QueryRow(
		"insert into user (username, hashed_password, last_password_change, created_at, updated_at) values ($1, $2, $3, $4, $5) returning id",
		u.Username, 
		u.HashedPassword, 
		u.LastPasswordChange, 
		u.CreatedAt,
		u.UpdatedAt,
	).Scan(&u.Id); err != nil {
		return nil, err
	}
	
	return u, nil
}

func (r *UserRepository) FindById(id int) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"select id, username, hashed_password, last_password_change, created_at, updated_at from notes where id = $1",
		id,
	).Scan(
		&u.Id,
		&u.Username,
		&u.HashedPassword,
		&u.LastPasswordChange,
		&u.CreatedAt,
		&u.UpdatedAt,
	); err != nil {
		return nil, err
	}
	
	return u, nil
}