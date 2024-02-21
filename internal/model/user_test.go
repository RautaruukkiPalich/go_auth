package model_test

import (
	"testing"

	"github.com/rautaruukkipalich/go_auth/internal/model"
	"github.com/stretchr/testify/assert"
)

type TestUser struct {
	name string
	u func() *model.User
	isValid bool
}

func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestUser(t)
	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.HashedPassword)
}

func TestUser_ValidateUsername(t *testing.T) {
	for _, tc := range getUsernameTestCases(t) {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})
	}
}

func TestUser_ValidatePassword(t *testing.T) {
	for _, tc := range getPasswordTestCases(t) {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})
	}
}

func TestUser_Validate(t *testing.T) {
	for _, tc := range getTestCases(t) {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})
	}
}

func getTestCases(t *testing.T) []TestUser {
	return append(getUsernameTestCases(t), getPasswordTestCases(t)...)
}

func getPasswordTestCases(t *testing.T) []TestUser {
	return []TestUser{
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
		{
			name: "password with letters and numbers",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = "aaa21331"
				return u	
			},
			isValid: true,
		},
		{
			name: "password with syblols",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = "aaa;bbb"
				return u	
			},
			isValid: false,
		},
	}
}

func getUsernameTestCases(t *testing.T) []TestUser {
	return []TestUser{
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
			name: "long username",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Username = "qwertyqwertyqwertyqwerty"
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
			name: "latin letters and numbers",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Username = "aaa1bbb"
				return u	
			},
			isValid: false,
		},
		{
			name: "cyrillic letters",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Username = "йцукен"
				return u	
			},
			isValid: false,
		},
	}
}
