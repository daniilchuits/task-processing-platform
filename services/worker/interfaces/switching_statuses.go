package interfaces

type Switcher interface {
	StatusProcessing(userId int, path string) error
}
