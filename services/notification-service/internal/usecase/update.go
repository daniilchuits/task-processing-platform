package usecase

import (
	"notifycation-service/internal/interfaces"
	"notifycation-service/internal/rabbitmq"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Msg struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
	Err    string `json:"err"`
}

type update struct {
	upd interfaces.Updater
}

func NewUpdate(upd interfaces.Updater) *update {
	return &update{upd: upd}
}

func (upd *update) UpdateWithDelivery(msg amqp.Delivery) error {

	domainMsg, err := rabbitmq.ToDomain(msg)
	if err != nil {
		return err
	}

	_, err = upd.upd.UpdateTasks(*domainMsg)
	if err != nil {
		return err
	}
	return nil
}
