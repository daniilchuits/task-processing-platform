package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type connManager struct {
	conn *amqp.Connection
}

func NewConnManager(conn *amqp.Connection) *connManager {
	return &connManager{conn: conn}
}
