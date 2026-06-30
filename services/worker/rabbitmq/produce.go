package rabbitmq

import (
	"encoding/json"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

type rabbitMqProducer struct {
	ch        *amqp091.Channel
	queueName string
}

func NewRabbitMqProducer(ch *amqp091.Channel, queueName string) *rabbitMqProducer {
	return &rabbitMqProducer{
		ch:        ch,
		queueName: queueName,
	}
}

type Msg struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
	Err    string `json:"err"`
}

func (prod *rabbitMqProducer) Produce(msg Msg) {

	msgBytes, err := json.Marshal(msg)
	if err != nil {
		log.Println("Error marshaling message:", err)
		return
	}

	if err = prod.ch.Publish(
		"",
		prod.queueName,
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        msgBytes,
		},
	); err != nil {
		log.Println("Error posting message:", err)
		return
	}
}
