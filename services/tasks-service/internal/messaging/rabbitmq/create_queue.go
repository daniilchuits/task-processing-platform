package rabbitmq

import (
	"github.com/rabbitmq/rabbitmq-amqp-go-client/pkg/rabbitmqamqp"
)

func (conn *conn) CreateQueue(queueName string) error {

	_, err := conn.connection.Management().DeclareQueue(
		conn.ctx,
		&rabbitmqamqp.QuorumQueueSpecification{Name: queueName},
	)
	return err
}
