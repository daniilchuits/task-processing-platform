package domain

import "errors"

var (

	// Validating
	ErrLogin            error = errors.New("Invalid length for login")
	ErrPassword         error = errors.New("Invalid length for password")
	ErrForbiddenSymbols error = errors.New("Forbidden symbols")

	// Hash
	ErrDuringHashing error = errors.New("Error during hashing")
	ErrWrongPassword error = errors.New("Wrong password")

	// JWT
	ErrCreatingJWT error = errors.New("Error creating jwt-token")

	// Register Usecase
	ErrUserExists error = errors.New("User already exists")

	// Login Usecase
	ErrUserNotExists error = errors.New("User not exists")

	// Internal Server error
	InternalServerError error = errors.New("Internal server error")
)
