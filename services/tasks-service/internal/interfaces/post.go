package interfaces

import "task-service/internal/domain"

type Poster interface {
	Insert(userID int, filename, filepath, typ string) (*domain.Task, error)
}
