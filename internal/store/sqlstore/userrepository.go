package sqlstore

import (
	"database/sql"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	model "github.com/rautaruukkipalich/go_auth/internal/model"
	"github.com/rautaruukkipalich/go_auth/internal/store"
)

type UserRepository struct {
	sqlstore *Store
}

func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	if err := u.BeforeCreate(); err != nil {
		return nil, err
	}
	
	utcNow := time.Now().UTC()

	u.CreatedAt = utcNow
	u.UpdatedAt = utcNow
	u.LastPasswordChange = utcNow

	err := r.sqlstore.db.QueryRow(
		"insert into users (username, hashed_password, last_password_change, created_at, updated_at) values ($1, $2, $3, $4, $5) returning id",
		u.Username, 
		u.HashedPassword, 
		u.CreatedAt, 
		u.UpdatedAt, 
		u.LastPasswordChange, 
	).Scan(&u.Id)
	
	if err != nil {
		return nil, err
	}
	
	return u, nil
}

func (r *UserRepository) FindById(id int) (*model.User, error) {
	u := &model.User{}
	if err := r.sqlstore.db.QueryRow(
		"select id, username, hashed_password, last_password_change, created_at, updated_at from users where id = $1",
		id,
	).Scan(
		&u.Id,
		&u.Username,
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

func (r *UserRepository) Find(username string) (*model.User, error) {
	u := &model.User{}
	if err := r.sqlstore.db.QueryRow(
		"select id, username, hashed_password, last_password_change, created_at, updated_at from users where username = $1",
		username,
	).Scan(
		&u.Id,
		&u.Username,
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

var jwtSecretKey = []byte("very-secret-key")

func (r *UserRepository) Auth(u *model.User) (string, error) {
	user, err := r.Find(u.Username)
	if err != nil {
		return "", store.ErrRecordNotFound
	}
	if !user.ComparePassword(u.Password){
		return "", store.ErrRecordNotFound
	}
	payload := jwt.MapClaims{
        "sub":  user.Id,
        "exp":  time.Now().Add(time.Minute * 20).Unix(),
    }

    // Создаем новый JWT-токен и подписываем его по алгоритму HS256
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

    signedToken, err := token.SignedString(jwtSecretKey)
    if err != nil {
        return "", errors.New("JWT token failed to signed")
    }
	return signedToken, nil
}

func (r *UserRepository) UpdatePassword(u *model.User) (*model.User, error) {
	return u, nil
}

func (r *UserRepository) UpdateUsername(u *model.User) (*model.User, error) {
	return u, nil
}