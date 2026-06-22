package domain

import "errors"

var (
	// Inteni
	ErrInternalServer error = errors.New("Internal server error")

	// Determine
	ErrInvalidExtension error = errors.New("Invalid extensionn")

	// Regexp
	ErrRegexp error = errors.New("Invalid filename")

	// Convert
	ErrConvUserId error = errors.New("Conv 'user_id' from request error")

	// Encoding
	ErrEncoding error = errors.New("Error encoding")

	// Insert
	ErrChecking         error = errors.New("Error during checking if note already exists")
	ErrExists           error = errors.New("That filename is already in the table")
	ErrInserting        error = errors.New("Error during inserting note")
	ErrDecodingFilename error = errors.New("Error decoding filename")

	// SelectAll
	ErrSelectingAll error = errors.New("Error selecting all tasks for user")

	// SelectOne
	ErrSelectingOne error = errors.New("Error selecting one task")
	ErrStrconvId    error = errors.New("Error converting task's id")
)
