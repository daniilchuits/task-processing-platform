package domain

import "errors"

var (

	// Validating
	ErrLogin            error = errors.New("Invalid length for login")
	ErrPassword         error = errors.New("Invalid length for password")
	ErrForbiddenSymbols error = errors.New("Forbidden symbols")

	ErrUserExists error = errors.New("User already exists")
)
