package rabbitmq

import (
	"github.com/rabbitmq/amqp091-go"
)

type connManager struct {
	conn      *amqp091.Connection
	queueName string
}

type rabbitMQConn interface {
	Consume() (<-chan amqp091.Delivery, error)
	CreateQueue() (*amqp091.Queue, error)
}

func NewConnManager(conn *amqp091.Connection, queueName string) rabbitMQConn {
	return &connManager{conn: conn, queueName: queueName}
}
