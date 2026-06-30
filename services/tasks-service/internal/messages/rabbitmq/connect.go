package rabbitmq

import (
	"time"

	"github.com/rabbitmq/amqp091-go"
)

func ConnectToRabbitMQ(brokerURI string) (*amqp091.Connection, error) {

	var err error

	for i := 0; i < 20; i++ {

		var err error
		conn, err := amqp091.Dial(brokerURI)
		if err == nil {
			return conn, nil
		}
		time.Sleep(time.Second)
	}
	return nil, err
}
