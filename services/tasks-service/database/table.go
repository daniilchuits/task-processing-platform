package database

import "database/sql"

type dbManager struct {
	db *sql.DB
}

func NewDbManager(db *sql.DB) dbManager {
	return dbManager{db: db}
}

func (manager dbManager) CreateTable() error {

	query := `
		CREATE TABLE IF NOT EXISTS tasks(
			id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
			user_id INT NOT NULL,
			filename TEXT NOT NULL,
			status TEXT NOT NULL,
			type TEXT NOT NULL,
			phrase_count_txt INT,
			main_colors_jpg TEXT,
			photo_resolution_jpg TEXT,
			audio_length_mp3 INT,
			num_of_lines_csv INT,
			pages_pdf INT,
			photo_in_doc_pdf BOOLEAN,
			size_after_unzip_zip INT,
			zip_files_zip TEXT
		);
	`

	_, err := manager.db.Exec(query)
	return err
}
