package teststore

import (
	"time"

	"github.com/rautaruukkipalich/go_auth/internal/model"
	"github.com/rautaruukkipalich/go_auth/internal/store"
)

type UserRepository struct {
	store *Store
	users map[string]*model.User
}

func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	if err := u.BeforeCreate(); err != nil {
		return nil, err
	}
	
	utcNow := time.Now().UTC()

	r.users[u.Username] = u
	u.Id = len(r.users)
	u.CreatedAt = utcNow
	u.UpdatedAt = utcNow
	u.LastPasswordChange = utcNow

	return r.users[u.Username], nil
}

func (r *UserRepository) Find(username string) (*model.User, error) {
	
	u := r.users[username]

	if u == nil {
		return nil, store.ErrRecordNotFound
	}

	return u, nil
}

func (r *UserRepository) FindById(id int) (*model.User, error) {
	
	var u *model.User
	
	for _, val := range r.users {
		if val.Id == id {
			return u, nil
		}
	}

	if u == nil {
		return nil, store.ErrRecordNotFound
	}

	return u, nil
}

func (r *UserRepository) Auth(*model.User) (string, error) {
	return "", nil
}
