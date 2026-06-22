package interfaces

import "task-service/internal/domain"

type NoteInserter interface {
	InsertNote(userId int, filename, filetype string) (*domain.Task, error)
}
