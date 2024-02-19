package model_test

import (
	"testing"

	"github.com/rautaruukkipalich/go_auth/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestUser(t)
	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.HashedPassword)
}

func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name string
		u func() *model.User
		isValid bool
	}{
		{
			name: "valid",
			u: func() *model.User {
				u := model.TestUser(t)
				return u	
			},
			isValid: true,
		},
		{
			name: "empty username",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Username = ""
				return u	
			},
			isValid: false,
		},
		{
			name: "username with spaces",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Username = "aaa bbb"
				return u	
			},
			isValid: false,
		},
		{
			name: "empty password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = ""
				return u	
			},
			isValid: false,
		},
		{
			name: "password with spaces",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = "aaa bbb"
				return u	
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})
	}
}
