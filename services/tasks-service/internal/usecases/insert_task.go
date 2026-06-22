package usecases

import (
	"log"
	"task-service/internal/domain"
	"task-service/internal/interfaces"
)

type InsertUsecase struct {
	Exists interfaces.CheckingExistence
	Insert interfaces.NoteInserter
}

func (insert *InsertUsecase) Exec(userId int, filename domain.Filename) (*domain.Task, error) {

	if err := domain.FilenameValidation(filename); err != nil {
		return nil, err
	}

	exists, err := insert.Exists.NoteExists(userId, filename.Name)
	if err != nil {
		log.Println(domain.ErrChecking, ":", err)
		return nil, domain.ErrChecking
	}
	if exists {
		return nil, domain.ErrExists
	}

	typ, err := domain.DetermineType(filename.Name)
	if err != nil {
		log.Println(err)
		return nil, domain.ErrInvalidExtension
	}

	task, err := insert.Insert.InsertNote(userId, filename.Name, typ)
	if err != nil { //
		log.Println("Error inserting:", err)
		return nil, domain.ErrInserting
	}

	return task, nil
}
