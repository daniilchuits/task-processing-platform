package usecases

import (
	"log"
	"task-service/internal/domain"
	"task-service/internal/interfaces"
)

type SelectOneTask struct {
	Selecter interfaces.Selecter
}

func (sel *SelectOneTask) Exec(taskId, userId int) (*domain.Task, error) {

	task, err := sel.Selecter.SelectOneTask(taskId, userId)
	if err != nil {
		log.Println(err)
		return nil, domain.ErrSelectingOne
	}
	return task, nil
}
