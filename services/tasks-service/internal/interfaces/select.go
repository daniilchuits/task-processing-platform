package interfaces

import "task-service/internal/domain"

type Selecter interface {
	SelectOneTask(taskId, userId int) (*domain.Task, error)
}
