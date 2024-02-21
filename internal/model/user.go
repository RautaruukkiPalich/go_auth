package model

import (
	"errors"
	"regexp"
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

var (
	ErrorContentUsername = errors.New("use latin letters only")
	ErrorContentPassword = errors.New("use letters and digits only")
	ErrorLenUsername = errors.New("username length must be between 3 and 20")
	ErrorLenPassword = errors.New("password length must be 3 synbols at least")
)

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
	// check username length
	matched, _ := regexp.MatchString(`^.{3,20}$`, username)
	if !matched {
		return ErrorLenUsername
	}
	// check username contains latin letters only
	matched, _ = regexp.MatchString(`^[A-z]+$`, username)
	if !matched {
		return ErrorContentUsername
	}

	return nil
}

func (u *User) ValidatePassword(password string) (err error) {
	
	// check password length
	matched, _ := regexp.MatchString(`^.{3,}$`, password)
	if !matched {
		return ErrorLenPassword
	}

	// check password contains leters and digits only
	matched, _ = regexp.MatchString(`^[\w\d]*$`, password)
	if !matched {
		return ErrorContentPassword
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
