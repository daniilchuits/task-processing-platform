package interfaces

type UserExistsInterface interface {
	UserExists(login string) (bool, error)
}
