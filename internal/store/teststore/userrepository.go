package teststore

import (
	"errors"
	"strings"
	"time"

	"github.com/rautaruukkipalich/go_auth/internal/model"
	"github.com/rautaruukkipalich/go_auth/internal/store"
	"github.com/rautaruukkipalich/go_auth/pkg/utils/jwt"
)

type UserRepository struct {
	store *Store
	users map[string]*model.User
}

func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	if err := u.BeforeCreate(); err != nil {
		return nil, err
	}
	
	if r.users[u.Username] != nil {
		return nil, errors.New("User already exists")
	}

	utcNow := time.Now().UTC()

	r.users[u.Username] = u
	u.Id = len(r.users)
	u.Slug = strings.ToLower(u.Username)
	u.CreatedAt = utcNow
	u.UpdatedAt = utcNow
	u.LastPasswordChange = utcNow

	return r.users[u.Username], nil
}

func (r *UserRepository) GetByUsername(username string) (*model.User, error) {
	
	u := r.users[username]

	if u == nil {
		return nil, store.ErrRecordNotFound
	}

	return u, nil
}

func (r *UserRepository) GetById(id int) (*model.User, error) {
	
	var u *model.User
	
	for _, val := range r.users {
		if val.Id == id {
			return val, nil
		}
	}

	if u == nil {
		return nil, store.ErrRecordNotFound
	}

	return u, nil
}

func (r *UserRepository) GetBySlug(slug string) (*model.User, error) {
	var u *model.User
	
	for _, val := range r.users {
		if val.Slug == slug {
			return val, nil
		}
	}

	if u == nil {
		return nil, store.ErrRecordNotFound
	}

	return u, nil
}

func (r *UserRepository) Auth(u *model.User) (string, error) {
	user, err := r.GetByUsername(u.Username)

	if err != nil {
		return "", store.ErrRecordNotFound
	}
	if !user.ComparePassword(u.Password){
		return "", store.ErrRecordNotFound
	}

	return jwt.EncodeJWTToken(user.Id)
}

func (r *UserRepository) SetPassword(u *model.User, password string) (error) {
	err := u.ValidatePassword(password)
	if err != nil {
		return errors.New(err.Error())
	}

	_, err = model.EncryptPassword(password)
	if err != nil {
		return errors.New(err.Error())
	}

	u.Password = password

	return nil
}

func (r *UserRepository) SetUsername(u *model.User, username string) (error) {
	err := u.ValidateUsername(username)
	if err != nil {
		return errors.New(err.Error())
	}

	slug := strings.ToLower(username)

	u.Username = username
	u.Slug = slug
	
	return nil
}