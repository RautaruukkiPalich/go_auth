package store

import "errors"

var (
	ErrRecordNotFound = errors.New("username or password is not valid")
	ErrInternalServerError = errors.New("internal server error")
)