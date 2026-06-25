package rabbitmq

import (
	"context"

	"github.com/rabbitmq/rabbitmq-amqp-go-client/pkg/rabbitmqamqp"
)

type MyPublisher struct {
	publisher *rabbitmqamqp.Publisher
	ctx       context.Context
}

func (conn *conn) NewPublisher(queueName string) (*MyPublisher, error) {
	publisher, err := conn.connection.NewPublisher(
		conn.ctx,
		&rabbitmqamqp.QueueAddress{Queue: queueName},
		nil,
	)
	if err != nil {
		return nil, err
	}
	return &MyPublisher{
		publisher: publisher,
		ctx:       conn.ctx,
	}, nil
}
