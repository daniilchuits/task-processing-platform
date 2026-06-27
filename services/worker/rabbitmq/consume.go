package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func Consume(ch *amqp.Channel, queueName string) (<-chan amqp.Delivery, error) {

	return ch.Consume(
		queueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
}
