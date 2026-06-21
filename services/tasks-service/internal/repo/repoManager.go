package repo

import "database/sql"

type repoManager struct {
	db *sql.DB
}

func NewRepoManager(db *sql.DB) *repoManager {
	return &repoManager{db: db}
}
