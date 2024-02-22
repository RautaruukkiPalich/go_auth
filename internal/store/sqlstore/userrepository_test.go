package sqlstore_test

import (
	"strings"
	"testing"

	model "github.com/rautaruukkipalich/go_auth/internal/model"
	"github.com/rautaruukkipalich/go_auth/internal/store"
	sqlstore "github.com/rautaruukkipalich/go_auth/internal/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseUrl)
	defer teardown("users")
	s := sqlstore.New(db)

	u, err := s.User().Create(model.TestUser(t))
	assert.NoError(t, err)
	assert.NotNil(t, u)

	_, err = s.User().Create(model.TestUser(t))
	assert.Error(t, err)
}

func TestUserRepository_FindById(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseUrl)
	defer teardown("users")
	s := sqlstore.New(db)

	// Test not existing user
	id := 0
	_, err := s.User().FindById(id)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u, _ := s.User().Create(model.TestUser(t))

	// Test existing user
	n, err := s.User().FindById(u.Id)
	assert.NoError(t, err)
	assert.NotNil(t, n)
}

func TestUserRepository_FindByUsername(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseUrl)
	defer teardown("users")
	s := sqlstore.New(db)

	// Test not existing user
	username := ""
	_, err := s.User().FindByUsername(username)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u, _ := s.User().Create(model.TestUser(t))

	// Test existing user
	n, err := s.User().FindByUsername(u.Username)
	assert.NoError(t, err)
	assert.NotNil(t, n)
}

func TestUserRepository_FindBySlug(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseUrl)
	defer teardown("users")
	s := sqlstore.New(db)

	// Test not existing user
	slug := ""
	_, err := s.User().FindBySlug(slug)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u, _ := s.User().Create(model.TestUser(t))

	// Test existing user
	n, err := s.User().FindBySlug(u.Slug)
	assert.NoError(t, err)
	assert.NotNil(t, n)
}

func TestUserRepository_Auth(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseUrl)
	defer teardown("users")
	s := sqlstore.New(db)
	
	u, _ := s.User().Create(model.TestUser(t))
	// Test valid user
	token, err := s.User().Auth(u)
	assert.NoError(t, err)
	assert.NotNil(t, token)

	// Test invalid password
	u.Password = "password123"
	token, err = s.User().Auth(u)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())
	assert.Len(t, token, 0)
}

func TestUserRepository_UpdatePassword(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseUrl)
	defer teardown("users")
	s := sqlstore.New(db)

	u, _ := s.User().Create(model.TestUser(t))
	new_password := "password123"
	// Test valid user
	err := s.User().UpdatePassword(u, new_password)
	assert.NoError(t, err)

	user, err := s.User().FindById(u.Id)
	assert.NoError(t, err)
	assert.True(t, user.ComparePassword(new_password))

	// Test invalid password
	new_password = ""
	err = s.User().UpdatePassword(u, new_password)
	assert.Error(t, err)

	user, err = s.User().FindById(u.Id)
	assert.NoError(t, err)
	assert.False(t, user.ComparePassword(new_password))
}

func TestUserRepository_UpdateUsername(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseUrl)
	defer teardown("users")
	s := sqlstore.New(db)

	u, _ := s.User().Create(model.TestUser(t))
	


	// Test valid user
	valid_username := "paSsword"
	valid_slug := strings.ToLower(valid_username)
	err := s.User().UpdateUsername(u, valid_username)
	assert.NoError(t, err)
	user, err := s.User().FindById(u.Id)
	assert.NoError(t, err)
	assert.Equal(t, user.Username, valid_username)
	assert.Equal(t, user.Slug, valid_slug)

	// Test invalid password
	invalid_username := ""
	invalid_slug := strings.ToLower(invalid_username)
	err = s.User().UpdateUsername(u, invalid_username)
	assert.Error(t, err)

	user, err = s.User().FindById(u.Id)
	assert.NoError(t, err)
	assert.NotEqual(t, user.Username, invalid_username)
	assert.NotEqual(t, user.Slug, invalid_slug)

	new_user, _ := s.User().Create(model.TestUser(t))
	err = s.User().UpdateUsername(new_user, valid_username)
	assert.Error(t, err)
}