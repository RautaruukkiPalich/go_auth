package model

import "time"

type User struct {
	Id                 int       `json:"id"`
	Username           string    `json:"username"`
	HashedPassword     string    `json:"hashed_password"`
	LastPasswordChange time.Time `json:"last_password_change"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}