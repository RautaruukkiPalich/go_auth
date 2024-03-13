package sqlstore

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/rautaruukkipalich/go_auth/internal/store"
)

const (
	createUser = `
		insert
		into users (username, slug, hashed_password, last_password_change, created_at, updated_at) 
		values ($1, $2, $3, $4, $5, $6)
		returning id
	`
	getUserByID = `
		select
		id, username, slug, hashed_password, last_password_change, created_at, updated_at 
		from users 
		where id = $1
	`
	getUserByUsername = `
		select
		id, username, slug, hashed_password, last_password_change, created_at, updated_at 
		from users 
		where username = $1
	`
	getUserBySlug = `
		select
		id, username, slug, hashed_password, last_password_change, created_at, updated_at 
		from users 
		where slug = $1
	`
	setPassword = `
		update users 
		set hashed_password=$1, last_password_change=$2, updated_at=$3 
		where id = $4
	`
	setUsername = `
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
		//take: username, slug, hashedPassword, lastPasswordChange, createdAt, updatedAt; 
		//return: userID 
		createUser *sql.Stmt
		//take: userID;
		//return: userID, username, slug, hashedPassword, lastPasswordChange, createdAt, updatedAt
		getUserByID *sql.Stmt
		//take: username;
		//return: userID, username, slug, hashedPassword, lastPasswordChange, createdAt, updatedAt
		getUserByUsername *sql.Stmt
		//take: slug;
		//return: userID, username, slug, hashedPassword, lastPasswordChange, createdAt, updatedAt
		getUserBySlug *sql.Stmt
		//take: hashedPassword, lastPasswordChange, updatedAt, userID
		setPassword *sql.Stmt
		//take: username, slug, updatedAt, userID
		setUsername *sql.Stmt
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

func (s *Store) User() store.UserRepositorier {
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
	createUser, err := db.Prepare(createUser)
	if err != nil {
		return nil, err
	}
	getUserByID, err := db.Prepare(getUserByID)
	if err != nil {
		return nil, err
	}
	getUserByUsername, err := db.Prepare(getUserByUsername)
	if err != nil {
		return nil, err
	}
	getUserBySlug, err := db.Prepare(getUserBySlug)
	if err != nil {
		return nil, err
	}
	setPassword, err := db.Prepare(setPassword)
	if err != nil {
		return nil, err
	}
	setUsername, err := db.Prepare(setUsername)
	if err != nil {
		return nil, err
	}
	return &UserStmts{
		createUser: createUser,
		getUserByID: getUserByID,
		getUserByUsername: getUserByUsername,
		getUserBySlug: getUserBySlug,
		setPassword: setPassword,
		setUsername: setUsername,
	}, nil
}