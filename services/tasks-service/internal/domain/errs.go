package domain

import "errors"

var (
	// Inteni
	ErrInternalServer error = errors.New("Internal server error")

	// Determine
	ErrInvalidExtension error = errors.New("Invalid extensionn")

	// CreateFile
	ErrFileExists error = errors.New("File with this name already exists")
	ErrCreating   error = errors.New("Error during creating file")
	ErrCopy       error = errors.New("Error to copy file")

	// Regexp
	ErrRegexp error = errors.New("Invalid filename")

	// Convert
	ErrConvUserId  error = errors.New("Conv 'user_id' from request error")
	ErrEmptyUserId error = errors.New("Empty 'user_id' in request")

	// Encoding
	ErrEncoding error = errors.New("Error encoding")

	// Insert
	ErrGetting                 error = errors.New("Error getting file")
	ErrDuringCheckingExistence error = errors.New("Error during checking existence")
	ErrExists                  error = errors.New("This user already has file with this name")
	ErrInserting               error = errors.New("Inserting error")

	// SelectAll
	ErrSelectingAll error = errors.New("Error selecting all tasks for user")

	// SelectOne
	ErrSelectingOne error = errors.New("Error selecting one task")
	ErrStrconvId    error = errors.New("Error converting task's id")

	// RabbitMQ
	ErrCreatingChannel error = errors.New("Error creating channel to rabbitmq")
	ErrCreatingQueue   error = errors.New("Error creating queue in rabbitmq")
	ErrSendingMessage  error = errors.New("Error sending message to rabbitmq")
)
