package rabbitmq

import (
	"time"

	"github.com/rabbitmq/amqp091-go"
)

func Conn(brokerURI string) (*amqp091.Connection, error) {

	var (
		conn *amqp091.Connection
		err  error
	)

	for i := 0; i < 20; i++ {

		var err error

		conn, err = amqp091.Dial(brokerURI)
		if err != nil {
			time.Sleep(time.Second)
			continue
		}
		break
	}
	return conn, err
}
