package rabbitmq

import (
	"log"
	"task-service/internal/domain"

	"github.com/Azure/go-amqp"
	rmq "github.com/rabbitmq/rabbitmq-amqp-go-client/pkg/rabbitmqamqp"
)

func (my *MyPublisher) Publish(msg []byte) error {

	message := rmq.NewMessage(msg)

	message.Header = &amqp.MessageHeader{
		Durable: true,
	}

	res, err := my.publisher.Publish(
		my.ctx,
		message,
	)
	if err != nil {
		log.Println("Error publishing message:", err)
		return domain.ErrPublishingMessageToRabbitMQ
	}

	switch res.Outcome.(type) {

	case *rmq.StateAccepted:
		return nil
	default:
		log.Panicf("Message not send")
	}
	return domain.ErrPublishingMessageToRabbitMQ
}
