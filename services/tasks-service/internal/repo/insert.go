package repo

import "task-service/internal/domain"

func (repo *repoManager) InsertNote(userId int, filename, filetype string) error {

	query := `
		INSERT INTO tasks (user_id, filename, status, type) VALUES
		($1,$2,$3,$4);
	`

	_, err := repo.db.Exec(query, userId, filename, domain.ProcessingStatus, filetype)
	// TODO
}
