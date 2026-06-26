package producer

import (
	"context"
	"encoding/json"
	"log"
	"task-service/internal/domain"
	"task-service/internal/messages/rabbitmq"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Message struct {
	Path     string `json:"path"`
	Filetype string `json:"filetype"`
}

func (conn *connManager) PublishMsg(
	ctx context.Context,
	path, filetype, queueName string,
) error {

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

	newMessage := Message{
		Path:     path,
		Filetype: filetype,
	}

	body, err := json.Marshal(newMessage)
	if err != nil {
		log.Println("Error marshaling newMessage:", err)
		return domain.ErrMarshaling
	}

	if err = ch.PublishWithContext(
		ctx,
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		},
	); err != nil {
		log.Println("Error publishing message to rabbitmq:", err)
		return domain.ErrSendingMessage
	}
	return nil
}
