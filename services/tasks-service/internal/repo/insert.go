package repo

import "task-service/internal/domain"

func (repo *repoManager) InsertNote(userId int, filename, filetype string) (*domain.Task, error) {

	query := `
		INSERT INTO tasks (user_id, filename, status, type) VALUES
		($1,$2,$3,$4)
		RETURNING id, user_id, filename, status, type
	`

	var task domain.Task

	if err := repo.db.QueryRow(query, userId, filename, domain.PendingStatus, filetype).Scan(
		&task.Id,
		&task.UserId,
		&task.Filename,
		&task.Status,
		&task.Typ,
	); err != nil {
		return nil, err
	}
	return &task, nil
}
