package rabbitmq

import "github.com/rabbitmq/amqp091-go"

func CreateQueueCosume(ch *amqp091.Channel, queueNameConsume string) (amqp091.Queue, error) {

	return ch.QueueDeclare(
		queueNameConsume,
		true,
		false,
		false,
		false,
		nil,
	)
}

func CreateQueueProduce(ch *amqp091.Channel, queueNameProd string) (amqp091.Queue, error) {
	return ch.QueueDeclare(
		queueNameProd,
		true,
		false,
		false,
		false,
		nil,
	)
}
