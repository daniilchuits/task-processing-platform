package interfaces

import "github.com/rabbitmq/amqp091-go"

type Consumer interface {
	Consume() (<-chan amqp091.Delivery, error)
	CreateQueue() (*amqp091.Queue, error)
}
