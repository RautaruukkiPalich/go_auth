package model

import (
	"regexp"
	"time"
	"log"

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
	// check username length
	matched, _ := regexp.MatchString(`^.{3,20}$`, username)
	if !matched {
		return ErrLenUsername
	}
	// check username contains latin letters only
	matched, _ = regexp.MatchString(`^[A-z]+$`, username)
	if !matched {
		return ErrContentUsername
	}

	return nil
}

func (u *User) ValidatePassword(password string) (err error) {
	
	// check password length (72 symbols - encrypt limit)
	matched, _ := regexp.MatchString(`^.{3,72}$`, password)
	if !matched {
		return ErrLenPassword
	}

	// check password contains leters and digits only
	matched, _ = regexp.MatchString(`^[\w\d]*$`, password)
	if !matched {
		return ErrContentPassword
	}

	return nil
}

func EncryptPassword(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Print(err)
		return "", ErrEncryptPassword
	}
// 
	return string(b), nil
}
