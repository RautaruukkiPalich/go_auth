package teststore_test

import (
	"strings"
	"testing"

	model "github.com/rautaruukkipalich/go_auth/internal/model"
	"github.com/rautaruukkipalich/go_auth/internal/store"
	"github.com/rautaruukkipalich/go_auth/internal/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()
	u, err := s.User().Create(model.TestUser(t))
	assert.NoError(t, err)
	assert.NotNil(t, u)
	_, err = s.User().Create(model.TestUser(t))
	assert.Error(t, err)
}

func TestUserRepository_FindById(t *testing.T) {
	s := teststore.New()

	// Test not existing user
	id := 12345
	_, err := s.User().FindById(id)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u, _ := s.User().Create(model.TestUser(t))

	// Test existing user
	n, err := s.User().FindById(u.Id)
	assert.NoError(t, err)
	assert.NotNil(t, n)
}

func TestUserRepository_FindByUsername(t *testing.T) {
	s := teststore.New()

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
	s := teststore.New()

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
	s := teststore.New()

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
	s := teststore.New()

	u, _ := s.User().Create(model.TestUser(t))
	new_password := "password123"
	// Test valid user
	err := s.User().UpdatePassword(u, new_password)
	assert.NoError(t, err)
	assert.Equal(t, u.Password, new_password)

	// Test invalid password
	new_password = ""
	err = s.User().UpdatePassword(u, new_password)
	assert.Error(t, err)
	assert.NotEqual(t, u.Password, new_password)
}

func TestUserRepository_UpdateUsername(t *testing.T) {
	s := teststore.New()

	u, _ := s.User().Create(model.TestUser(t))
	new_username := "paSsword"
	slug := strings.ToLower(new_username)
	// Test valid user
	err := s.User().UpdateUsername(u, new_username)
	assert.NoError(t, err)
	assert.Equal(t, u.Username, new_username)
	assert.Equal(t, u.Slug, slug)

	// Test invalid password
	new_username = ""
	slug = strings.ToLower(new_username)
	err = s.User().UpdateUsername(u, new_username)
	assert.Error(t, err)
	assert.NotEqual(t, u.Username, new_username)
	assert.NotEqual(t, u.Slug, slug)
}
