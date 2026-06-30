package repo

import "database/sql"

type notifyRepo struct {
	db *sql.DB
}

func NewNotifyRepo(db *sql.DB) *notifyRepo {
	return &notifyRepo{db: db}
}
