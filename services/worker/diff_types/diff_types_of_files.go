package difftypes

import (
	"encoding/json"
	"log"
	"worker/domain"
	"worker/interfaces"
	"worker/txt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Message struct {
	UserId   int    `json:"user_id"`
	Path     string `json:"path"`
	Filetype string `json:"filetype"`
}

type ultimateStruct struct {
	txt txt.TxtUpdate
}

func NewUltimateStruct(
	updater interfaces.TxtUpdater,
	switcher interfaces.Switcher,
) *ultimateStruct {
	return &ultimateStruct{txt: txt.TxtUpdate{
		Txt:      updater,
		Switcher: switcher,
	}}
}

func (ult *ultimateStruct) DistributeFiles(delivery amqp.Delivery) error {

	var message Message
	if err := json.Unmarshal(delivery.Body, &message); err != nil {
		return err
	}

	log.Println("Received message:", message)

	switch message.Filetype {

	case domain.TxtExtension:
		if err := ult.txt.Work(message.UserId, message.Path); err != nil {
			log.Println("Error operating txt:", err)
		} else {
			log.Println("Task done")
		}

		// case domain.CsvExtension:
		// 	if err := workCsv(message.Path); err != nil {
		// 		log.Println("Error operating csv:", err)
		// 	}
		// case domain.JpgExtension:
		// 	if err := workJpg(message.Path); err != nil {
		// 		log.Println("Error operating jpg:", err)
		// 	}
		// case domain.Mp3Extension:
		// 	if err := workMp3(message.Path); err != nil {
		// 		log.Println("Error operating mp3:", err)
		// 	}
		// case domain.PdfExtension:
		// 	if err := workPdf(message.Path); err != nil {
		// 		log.Println("Error operating pdf:", err)
		// 	}
		// case domain.ZipExtension:
		// 	if err := workZip(message.Path); err != nil {
		// 		log.Println("Error operating zip:", err)
		// 	}
	}
	return nil
}
