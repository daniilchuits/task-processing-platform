package rabbitmq

import (
	"encoding/json"
	"notifycation-service/internal/domain"

	"github.com/rabbitmq/amqp091-go"
)

func ToDomain(msg amqp091.Delivery) (*domain.Msg, error) {

	var msgDom domain.Msg
	if err := json.Unmarshal(msg.Body, &msgDom); err != nil {
		return nil, err
	}

	return &msgDom, nil
}
