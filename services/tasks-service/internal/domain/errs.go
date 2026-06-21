package domain

import "errors"

var (
	// Determine
	ErrInvalidExtension error = errors.New("Invalid extensionn")

	// Insert
	ErrChecking  error = errors.New("Error during checking if note already exists")
	ErrExists    error = errors.New("That filename is already in the table")
	ErrInserting error = errors.New("Error during inserting note")
)
