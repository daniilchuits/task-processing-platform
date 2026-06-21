package interfaces

type NoteInserter interface {
	InsertNote(userId int, filename, filetype string) error
}
