package interfaces

import (
	"worker/rabbitmq"
)

type Producer interface {
	Produce(msg rabbitmq.Msg)
}
