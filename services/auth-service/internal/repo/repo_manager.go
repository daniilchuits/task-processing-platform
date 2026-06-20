package repo

import "database/sql"

type repoManager struct {
	repo *sql.DB
}

func NewRepoManager(repo *sql.DB) *repoManager {
	return &repoManager{repo: repo}
}
