package model

import "errors"

var (
	ErrContentUsername = errors.New("use latin letters only")
	ErrContentPassword = errors.New("use letters and digits only")
	ErrLenUsername     = errors.New("username length must be between 3 and 20")
	ErrLenPassword     = errors.New("password length must be between 3 and 72")
	ErrEncryptPassword = errors.New("internal server error")
)
