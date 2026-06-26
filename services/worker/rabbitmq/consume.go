package rabbitmq

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Consume(ctx context.Context, ch *amqp.Channel, queueName string) (<-chan amqp.Delivery, error) {

	return ch.ConsumeWithContext(
		ctx,
		queueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
}
