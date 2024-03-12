package sqlstore

import (
	"database/sql"
	"errors"
	"log"
	"strings"
	"time"

	model "github.com/rautaruukkipalich/go_auth/internal/model"
	"github.com/rautaruukkipalich/go_auth/internal/store"
	"github.com/rautaruukkipalich/go_auth/internal/utils"
)

func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	if err := u.BeforeCreate(); err != nil {
		return nil, err
	}
	
	utcNow := time.Now().UTC()

	u.Slug = strings.ToLower(u.Username)
	u.CreatedAt = utcNow
	u.UpdatedAt = utcNow
	u.LastPasswordChange = utcNow

	err := r.stmts.insertUser.QueryRow(
		u.Username, 
		u.Slug,
		u.HashedPassword, 
		u.LastPasswordChange, 
		u.CreatedAt, 
		u.UpdatedAt, 
	).Scan(&u.Id)
	
	if err != nil {
		log.Printf("User: %v, err: %v; type: %T", u.Username, err, err)
		return nil, store.ErrInternalServerError
	}
	
	return u, nil
}

func (r *UserRepository) FindById(id int) (*model.User, error) {
	u := &model.User{}

	err := r.stmts.findUserByID.QueryRow(
		id, 
	).Scan(
		&u.Id,
		&u.Username,
		&u.Slug,
		&u.HashedPassword,
		&u.LastPasswordChange,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		} 
		log.Printf("User: %v, err: %v; type: %T", u.Username, err, err)
		return nil, store.ErrInternalServerError
	}
	
	return u, nil
}

func (r *UserRepository) FindByUsername(username string) (*model.User, error) {
	u := &model.User{}

	err := r.stmts.findUserByUsername.QueryRow(
		username, 
	).Scan(
		&u.Id,
		&u.Username,
		&u.Slug,
		&u.HashedPassword,
		&u.LastPasswordChange,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		} 
		log.Printf("User: %v, err: %v; type: %T", u.Username, err, err)
		return nil, store.ErrInternalServerError
	}
	
	return u, nil
}

func (r *UserRepository) FindBySlug(slug string) (*model.User, error) {
	u := &model.User{}

	err := r.stmts.findUserBySlug.QueryRow(
		slug, 
	).Scan(
		&u.Id,
		&u.Username,
		&u.Slug,
		&u.HashedPassword,
		&u.LastPasswordChange,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		} 
		log.Printf("User: %v, err: %v; type: %T", u.Username, err, err)
		return nil, store.ErrInternalServerError
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
		return err
	}

	hashedPassword, err := model.EncryptPassword(password)
	if err != nil {
		return err
	}

	_, err = r.stmts.updatePassword.Exec(
		hashedPassword,
		time.Now().UTC(),
		time.Now().UTC(),
		u.Id,
	)

	if err != nil {
		log.Printf("User: %v, err: %v; type: %T", u.Username, err, err)
		return store.ErrInternalServerError
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
				log.Printf("User: %v, err: %v; type: %T", u.Username, err, err)
				return store.ErrInternalServerError
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
		log.Printf("User: %v, err: %v; type: %T", u.Username, err, err)
		return store.ErrInternalServerError
	}

	return nil
}
