package domain

import "errors"

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrNotFound           = errors.New("register not found")
)
