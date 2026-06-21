package domain

import "errors"

var (
	ErrNoBearerPrefix       error = errors.New("No 'bearer' prefix")
	ErrWrongTokenMethod     error = errors.New("Wrong signing method")
	ErrTokenInvalid         error = errors.New("Token is already invalid")
	ErrGettingClaimsFromJWT error = errors.New("Error getting claims from jwt")
	ErrNotFound             error = errors.New("'user_id' not found in jwt")
)
