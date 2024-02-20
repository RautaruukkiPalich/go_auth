package sqlstore

import (
	"database/sql"
	"errors"
	"strings"
	"time"

	model "github.com/rautaruukkipalich/go_auth/internal/model"
	"github.com/rautaruukkipalich/go_auth/internal/store"
	"github.com/rautaruukkipalich/go_auth/internal/utils"
)

type UserRepository struct {
	sqlstore *Store
}


func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	if err := u.BeforeCreate(); err != nil {
		return nil, err
	}
	
	utcNow := time.Now().UTC()

	u.Slug = strings.ToLower(u.Username)
	u.CreatedAt = utcNow
	u.UpdatedAt = utcNow
	u.LastPasswordChange = utcNow

	err := r.sqlstore.db.QueryRow(
		"insert into users (username, slug, hashed_password, last_password_change, created_at, updated_at) values ($1, $2, $3, $4, $5, $6) returning id",
		u.Username, 
		u.Slug,
		u.HashedPassword, 
		u.LastPasswordChange, 
		u.CreatedAt, 
		u.UpdatedAt, 
	).Scan(&u.Id)
	
	if err != nil {
		return nil, err
	}
	
	return u, nil
}

func (r *UserRepository) FindById(id int) (*model.User, error) {
	u := &model.User{}
	if err := r.sqlstore.db.QueryRow(
		"select id, username, slug, hashed_password, last_password_change, created_at, updated_at from users where id = $1",
		id,
	).Scan(
		&u.Id,
		&u.Username,
		&u.Slug,
		&u.HashedPassword,
		&u.LastPasswordChange,
		&u.CreatedAt,
		&u.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		} 
		return nil, err
	}
	
	return u, nil
}

func (r *UserRepository) FindByUsername(username string) (*model.User, error) {
	u := &model.User{}
	if err := r.sqlstore.db.QueryRow(
		"select id, username, slug, hashed_password, last_password_change, created_at, updated_at from users where username = $1",
		username,
	).Scan(
		&u.Id,
		&u.Username,
		&u.Slug,
		&u.HashedPassword,
		&u.LastPasswordChange,
		&u.CreatedAt,
		&u.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		} 
		return nil, err
	}
	
	return u, nil
}

func (r *UserRepository) FindBySlug(slug string) (*model.User, error) {
	u := &model.User{}
	if err := r.sqlstore.db.QueryRow(
		"select id, username, slug, hashed_password, last_password_change, created_at, updated_at from users where slug = $1",
		slug,
	).Scan(
		&u.Id,
		&u.Username,
		&u.Slug,
		&u.HashedPassword,
		&u.LastPasswordChange,
		&u.CreatedAt,
		&u.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		} 
		return nil, err
	}
	
	return u, nil
}

func (r *UserRepository) Auth(u *model.User) (string, error) {

	user, err := r.FindByUsername(u.Username)

	if err != nil {
		return "", store.ErrRecordNotFound
	}
	if !user.ComparePassword(u.Password){
		return "", store.ErrRecordNotFound
	}

	return utils.EncodeJWTToken(user)
}

func (r *UserRepository) UpdatePassword(u *model.User, password string) (error) {
	err := u.ValidatePassword(password)
	if err != nil {
		return errors.New(err.Error())
	}

	hashedPassword, err := model.EncryptPassword(password)
	if err != nil {
		return errors.New(err.Error())
	}

	_, err = r.sqlstore.db.Exec(
		"update users set hashed_password=$1, last_password_change=$2, updated_at=$3 where id = $4",
		hashedPassword,
		time.Now().UTC(),
		time.Now().UTC(),
		u.Id,
	)
	if err != nil {
		return errors.New(err.Error())
	}	

	return nil
}

func (r *UserRepository) UpdateUsername(u *model.User, username string) (error) {
	err := u.ValidateUsername(username)
	if err != nil {
		return errors.New(err.Error())
	}

	slug := strings.ToLower(username)

	if u.Slug != slug {
		_, err = r.FindBySlug(slug)
		if err != nil {
			switch err{
			case store.ErrRecordNotFound:
				break
			default:
				return errors.New(err.Error())
			}
		}
	}
	
	_, err = r.sqlstore.db.Exec(
		"update users set username=$1, slug=$2, updated_at=$3 where id = $4",
		username,
		slug,
		time.Now().UTC(),
		u.Id,
	)
	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}
