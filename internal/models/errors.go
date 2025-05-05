package models

import "errors"

var (
	ErrNoRecord = errors.New("models: no matching record found")

	ErrInvalid = errors.New("models: invalid credentials")

	ErrDuplicateEmail = errors.New("models: duplicate email")

	ErrInactiveAccount = errors.New("models: inactive account")

	ErrInvalidCredentials = errors.New("models: invalid credentials")
)
