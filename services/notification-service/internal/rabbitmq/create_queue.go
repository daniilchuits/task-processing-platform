package rabbitmq

import "github.com/rabbitmq/amqp091-go"

func (manager *connManager) CreateQueue() (*amqp091.Queue, error) {

	ch, err := manager.conn.Channel()
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		manager.queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	return &q, err
}
