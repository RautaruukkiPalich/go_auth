package sqlstore

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/rautaruukkipalich/go_auth/internal/store"
)

const (
	insertUser = `
		insert
		into users (username, slug, hashed_password, last_password_change, created_at, updated_at) 
		values ($1, $2, $3, $4, $5, $6)
		returning id
	`
	findUserByID = `
		select
		id, username, slug, hashed_password, last_password_change, created_at, updated_at 
		from users 
		where id = $1
	`
	findUserByUsername = `
		select
		id, username, slug, hashed_password, last_password_change, created_at, updated_at 
		from users 
		where username = $1
	`
	findUserBySlug = `
		select
		id, username, slug, hashed_password, last_password_change, created_at, updated_at 
		from users 
		where slug = $1
	`
	updatePassword = `
		update users 
		set hashed_password=$1, last_password_change=$2, updated_at=$3 
		where id = $4
	`
	updateUsername = `
		update users 
		set username=$1, slug=$2, updated_at=$3 
		where id = $4
	`
)

type (
	Store struct {
		db             *sql.DB
		userRepository *UserRepository
	}
	
	UserRepository struct {
		sqlstore *Store
		stmts *UserStmts
	}
	
	UserStmts struct {
		insertUser *sql.Stmt
		findUserByID *sql.Stmt
		findUserByUsername *sql.Stmt
		findUserBySlug *sql.Stmt
		updatePassword *sql.Stmt
		updateUsername *sql.Stmt
	}

)

func New(db *sql.DB) (*Store, error) {
	store := &Store{
		db: db,
	}
	
	userRepository, err := NewUserRepository(store)
	if err != nil {
		return nil, err
	}
	store.userRepository = userRepository

	return store, nil
}

func (s *Store) User() store.UserRepository {
	return s.userRepository
}

func NewUserRepository (s *Store) (*UserRepository, error) {
	stmts, err := prepareStatements(s.db)
	if err != nil {
		return nil, err
	}

	return &UserRepository{
		sqlstore: s,
		stmts: stmts,
	}, nil
}

func prepareStatements(db *sql.DB) (*UserStmts, error) {
	insertUser, err := db.Prepare(insertUser)
	if err != nil {
		return nil, err
	}
	findUserByID, err := db.Prepare(findUserByID)
	if err != nil {
		return nil, err
	}
	findUserByUsername, err := db.Prepare(findUserByUsername)
	if err != nil {
		return nil, err
	}
	findUserBySlug, err := db.Prepare(findUserBySlug)
	if err != nil {
		return nil, err
	}
	updatePassword, err := db.Prepare(updatePassword)
	if err != nil {
		return nil, err
	}
	updateUsername, err := db.Prepare(updateUsername)
	if err != nil {
		return nil, err
	}
	return &UserStmts{
		insertUser: insertUser,
		findUserByID: findUserByID,
		findUserByUsername: findUserByUsername,
		findUserBySlug: findUserBySlug,
		updatePassword: updatePassword,
		updateUsername: updateUsername,
	}, nil
}