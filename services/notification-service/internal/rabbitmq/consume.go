package rabbitmq

import "github.com/rabbitmq/amqp091-go"

func (manager *connManager) Consume() (<-chan amqp091.Delivery, error) {

	ch, err := manager.conn.Channel()
	if err != nil {
		return nil, err
	}

	return ch.Consume(
		manager.queueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
}
