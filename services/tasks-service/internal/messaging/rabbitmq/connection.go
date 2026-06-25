package rabbitmq

import (
	"context"

	rmq "github.com/rabbitmq/rabbitmq-amqp-go-client/pkg/rabbitmqamqp"
)

type conn struct {
	connection *rmq.AmqpConnection
	ctx        context.Context
}

func NewConn(brokerURI string, ctx context.Context) (*conn, error) {

	env := rmq.NewEnvironment(brokerURI, nil)
	c, err := env.NewConnection(ctx)
	if err != nil {
		return nil, err
	}
	return &conn{connection: c, ctx: ctx}, nil
}
