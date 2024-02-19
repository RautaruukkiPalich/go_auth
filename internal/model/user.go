package model

import (
	"errors"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id                 int       `json:"id"`
	Username           string    `json:"username"`
	Password     	   string    `json:"password,omitempty"`
	HashedPassword     string    `json:"hashed_password"`
	LastPasswordChange time.Time `json:"last_password_change"`
	CreatedAt          time.Time `json:"created_at,omitempty"`
	UpdatedAt          time.Time `json:"updated_at,omitempty"`
}

func (u *User) BeforeCreate() error {
	if err := u.Validate(); err != nil {
		return err
	}
	enc, err := EncryptPassword(u.Password)
	if err != nil {
		return err
	}
	u.HashedPassword = enc

	return nil
}

func (u *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.HashedPassword), []byte(password)) == nil
}

func (u *User) Sanitize() {
	u.Password = ""
}

func (u *User) Validate() error {
	if err := u.ValidateUsername(u.Username); err != nil {
		return err
	}
	if err := u.ValidatePassword(u.Password); err != nil {
		return err
	}
	return nil
}

func (u *User) ValidateUsername(username string) (err error) {
	if len(username) < 1 {
		return errors.New("username must be longer then one character")
	}

	if len(username) > 100 {
		return errors.New("username can not be longer then 100 character")
	}

	splitUsername := strings.Split(username, " ")
	if len(splitUsername) > 1 {
		return errors.New("username can not contains spaces")
	}


	// Valid with regexp
	// Valid unique

	return nil
}

func (u *User) ValidatePassword(password string) (err error) {
	if len(password) < 1 {
		return errors.New("password must be longer then one character")
	}
	splitPassword := strings.Split(password, " ")
	if len(splitPassword) > 1 {
		return errors.New("password can not contains spaces")
	}
	// Valid with regexp

	return nil
}

func EncryptPassword(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
