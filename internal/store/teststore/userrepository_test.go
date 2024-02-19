package teststore_test

import (
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