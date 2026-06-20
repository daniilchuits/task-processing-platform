package database

import "database/sql"

type dbManager struct {
	db *sql.DB
}

func NewDbManager(db *sql.DB) *dbManager {
	return &dbManager{db: db}
}

func (manager *dbManager) Create() error {

	query := `
		CREATE IF NOT EXISTS TABLE users(
			id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
			login TEXT NOT NULL,
			hashed_password TEXT NOT NULL
		);
	`

	_, err := manager.db.Exec(query)
	return err
}
