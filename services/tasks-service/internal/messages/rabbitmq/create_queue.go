package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func CreateQueue(ch *amqp.Channel, queueName string) (amqp.Queue, error) {
	return ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
}
