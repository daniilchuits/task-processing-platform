package interfaces

import "task-service/internal/domain"

type SelecterAll interface {
	SelectAllTasks(userId int) (*[]domain.Task, error)
}
