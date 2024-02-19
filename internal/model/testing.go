package model

import "testing"

func TestUser(t *testing.T) *User {
	t.Helper()
	return &User{
		Username: "testuser",
		Password: "testpassword",
	}
}