package repo

import (
	"errors"
	"fmt"
	"task-service/internal/domain"
)

var PostingErr error = errors.New("Error posing")

func (manager *repoManager) Insert(userID int, filename, filepath, typ string) (*domain.Task, error) {

	query := `
		INSERT INTO tasks (user_id,filename,filepath,status,type) VALUES
		($1,$2,$3,$4,$5)
		RETURNING id,user_id,filename,filepath,status,type
	`

	var task domain.Task

	if err := manager.db.QueryRow(
		query,
		userID,
		filename,
		filepath,
		domain.PendingStatus,
		typ,
	).Scan(
		&task.Id,
		&task.UserId,
		&task.Filename,
		&task.Filepath,
		&task.Status,
		&task.Typ,
	); err != nil {
		return nil, fmt.Errorf("%w: %w", PostingErr, err)
	}
	return &task, nil
}
