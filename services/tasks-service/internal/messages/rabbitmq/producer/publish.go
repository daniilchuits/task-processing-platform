package producer

import (
	"context"
	"log"
	"task-service/internal/domain"
	"task-service/internal/messages/rabbitmq"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (conn *connManager) PublishMsg(ctx context.Context, msg, queueName string) error {

	ch, err := rabbitmq.CreateChannel(conn.conn)
	if err != nil {
		log.Println("Creating channel in rabbitmq err:", err)
		return domain.ErrCreatingChannel
	}
	defer ch.Close()

	q, err := rabbitmq.CreateQueue(ch, queueName)
	if err != nil {
		log.Println("Creating queue in rabbitmq err:", err)
		return domain.ErrCreatingQueue
	}

	if err = ch.PublishWithContext(
		ctx,
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		},
	); err != nil {
		log.Println("Error publishing message to rabbitmq:", err)
		return domain.ErrSendingMessage
	}
	return nil
}
