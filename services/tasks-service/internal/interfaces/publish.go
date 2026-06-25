package interfaces

type Publisher interface {
	Publish(msg []byte) error
}
