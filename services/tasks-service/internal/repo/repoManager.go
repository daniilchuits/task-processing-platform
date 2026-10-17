package repo

import "database/sql"

type repoManager struct {
	db db
}

type row interface {
	Scan(dest ...any) error
}

type db interface {
	Query(query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...any) row
}

func (repo repoManager) QueryRow(query string, args ...any) row {
	return repo.db.QueryRow(query, args...)
}

func NewRepoManager(db db) *repoManager {
	return &repoManager{db: db}
}
