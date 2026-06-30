package rabbitmq

import (
	"notifycation-service/internal/interfaces"

	"github.com/rabbitmq/amqp091-go"
)

type connManager struct {
	conn      *amqp091.Connection
	queueName string
}

func NewConnManager(conn *amqp091.Connection, queueName string) interfaces.Consumer {
	return &connManager{conn: conn, queueName: queueName}
}
