package interfaces

type Checker interface {
	Check(userID int, filename string) (bool, error)
}
