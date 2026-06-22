package usecases

import (
	"task-service/internal/domain"
	"task-service/internal/interfaces"
)

type SelectAllTasksUsecase struct {
	Selecter interfaces.SelecterAll
}

func (sel *SelectAllTasksUsecase) Exec(userId int) (*[]domain.Task, error) {
	tasks, err := sel.Selecter.SelectAllTasks(userId)
	if err != nil {
		return nil, domain.ErrSelectingAll
	}
	return tasks, nil
}
