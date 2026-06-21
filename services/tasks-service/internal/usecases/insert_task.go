package usecases

import (
	"log"
	"task-service/internal/domain"
	"task-service/internal/interfaces"
)

type insertUsecase struct {
	exists interfaces.CheckingExistence
	insert interfaces.NoteInserter
}

func (insert *insertUsecase) Exec(userId int, filename domain.Filename) error {

	exists, err := insert.exists.NoteExists(userId, filename.Name)
	if err != nil {
		log.Println(domain.ErrChecking, ":", err)
		return domain.ErrChecking
	}
	if exists {
		return domain.ErrExists
	}

	typ, err := domain.DetermineType(filename.Name)
	if err != nil {
		log.Println(err)
		return domain.ErrInvalidExtension
	}

	if err = insert.insert.InsertNote(userId, filename.Name, typ); err != nil { //
		log.Println("Error inserting:", err)
		return domain.ErrInserting
	}

	return nil
}
