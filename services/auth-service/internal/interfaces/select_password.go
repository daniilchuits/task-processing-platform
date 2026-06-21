package interfaces

type SelPasswordInterface interface {
	SelectPassword(login string) (string, int, error)
}
