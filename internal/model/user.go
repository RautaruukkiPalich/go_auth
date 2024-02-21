package model

import (
	"errors"
	"regexp"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id                 int       `json:"id"`
	Username           string    `json:"username"`
	Slug 			   string	 `json:"slug,omitempty"`
	Password     	   string    `json:"password,omitempty"`
	HashedPassword     string    `json:"hashed_password,omitempty"`
	LastPasswordChange time.Time `json:"last_password_change"`
	CreatedAt          time.Time `json:"created_at"`
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
	if len(username) < 3 {
		return errors.New("username must be longer then 3 characters")
	}

	if len(username) > 20 {
		return errors.New("username can not be longer then 20 characters")
	}

	// Valid with regexp
	matched, err := regexp.MatchString(`^[A-z]+$`, username)
	if err != nil {
		return errors.New("server error")
	}
	if !matched {
		return errors.New("use latin letters only")
	}

	return nil
}

func (u *User) ValidatePassword(password string) (err error) {
	if len(password) < 4 {
		return errors.New("password must be longer then one character")
	}

	splitPassword := strings.Split(password, " ")
	if len(splitPassword) > 1 {
		return errors.New("password can not contains spaces")
	}

	// Valid with regexp
	matched, err := regexp.MatchString(`^[\w\d]*$`, password)
	if err != nil {
		return errors.New("server error")
	}
	if !matched {
		return errors.New("use letters and digits only")
	}

	return nil
}

func EncryptPassword(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
