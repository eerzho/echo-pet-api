package exception

import "errors"

var (
	ErrUnauthorized = errors.New("unauthorized")
	ErrInvalidLogin = errors.New("invalid login")
)
