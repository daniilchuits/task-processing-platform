package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func CreateChannel(conn *amqp.Connection) (*amqp.Channel, error) {
	return conn.Channel()
}
