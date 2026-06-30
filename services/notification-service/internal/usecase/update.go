package usecase

import (
	"log"
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

	log.Println("Notification 1")
	domainMsg, err := rabbitmq.ToDomain(msg)
	if err != nil {
		return err
	}
	log.Println("Notification 2")

	err = upd.upd.UpdateTasks(*domainMsg)
	if err != nil {
		return err
	}
	log.Println("Notification 3")
	return nil
}
