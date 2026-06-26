package rabbitmq

import "github.com/rabbitmq/amqp091-go"

func CreateQueue(ch *amqp091.Channel, queueName string) (amqp091.Queue, error) {

	return ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		amqp091.Table{
			amqp091.QueueTypeArg: amqp091.QueueTypeQuorum,
		},
	)
}
