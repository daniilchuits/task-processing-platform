package rabbitmq

import (
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Connect(brokerURI string) (*amqp.Connection, error) {

	var (
		err error
	)

	for i := 0; i < 20; i++ {

		time.Sleep(time.Second)

		var err error
		conn, err := amqp.Dial(brokerURI)
		if err == nil {
			return conn, nil
		}

	}
	return nil, err
}
