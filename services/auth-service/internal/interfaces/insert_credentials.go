package interfaces

import "auth/internal/domain"

type InsertCredentialsInterface interface {
	InsertCredentials(cred domain.Credentials) (*domain.Credentials, error)
}
