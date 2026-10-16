package repo

import (
	"auth/internal/domain"
	"database/sql"
)

type repoManagerReal struct {
	repo *sql.DB
}

type repoManagerMock struct{}

type repoManagerMockWErrUserExists struct {
	err    error
	exists bool
}

type repoManager interface {
	InsertCredentials(credentials domain.Credentials) (*domain.Credentials, error)
	SelectPassword(login string) (string, int, error)
	UserExists(login string) (bool, error)
}

func NewRepoManager(repo *sql.DB) repoManager {
	return repoManagerReal{repo: repo}
}
