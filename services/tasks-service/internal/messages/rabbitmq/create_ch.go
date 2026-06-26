package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func createChannel(conn *amqp.Connection) (*amqp.Channel, error) {
	return conn.Channel()
}
