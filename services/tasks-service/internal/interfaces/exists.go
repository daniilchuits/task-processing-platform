package interfaces

type CheckingExistence interface {
	NoteExists(userId int, filename string) (bool, error)
}
