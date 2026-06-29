package difftypes

import (
	"encoding/json"
	"log"
	"time"
	"worker/csv"
	"worker/domain"
	"worker/interfaces"
	"worker/jpg"
	"worker/mp3"
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
	jpg jpg.JpegUpdate
	mp3 mp3.AudioUpdate
	csv csv.CSVUpdater
}

func NewUltimateStruct(
	updaterTxt interfaces.TxtUpdater,
	switcher interfaces.Switcher,
	updaterJPG interfaces.JPGUpdater,
	updaterMP3 interfaces.MP3Updater,
	updaterCsv interfaces.CSVUpdater,
) *ultimateStruct {
	return &ultimateStruct{txt: txt.TxtUpdate{
		Txt:      updaterTxt,
		Switcher: switcher,
	}, jpg: jpg.JpegUpdate{
		Switcher: switcher,
		Update:   updaterJPG,
	}, mp3: mp3.AudioUpdate{
		Switcher: switcher,
		Updater:  updaterMP3,
	}, csv: csv.CSVUpdater{
		Switcher: switcher,
		Update:   updaterCsv,
	},
	}
}

func (ult *ultimateStruct) DistributeFiles(delivery amqp.Delivery) error {

	var message Message
	if err := json.Unmarshal(delivery.Body, &message); err != nil {
		return err
	}

	switch message.Filetype {

	case domain.TxtExtension:
		if err := ult.txt.Work(message.UserId, message.Path); err != nil {
			log.Println("Error operating txt:", err)
		}
		time.Sleep(2 * time.Second)
	case domain.JpgExtension:
		if err := ult.jpg.Work(message.UserId, message.Path); err != nil {
			log.Println("Error operating jpg:", err)
		}
		time.Sleep(5 * time.Second)
	case domain.Mp3Extension:
		if err := ult.mp3.Work(message.UserId, message.Path); err != nil {
			log.Println("Error operating mp3:", err)
		}
		time.Sleep(3 * time.Second)

	case domain.CsvExtension:
		if err := ult.csv.Work(message.UserId, message.Path); err != nil {
			log.Println("Error operating csv:", err)
		}
		time.Sleep(time.Second)
	case domain.ZipExtension:
		if err := workZip(message.Path); err != nil {
			log.Println("Error operating zip:", err)
		}
		time.Sleep()

		// case domain.PdfExtension: // worker для pdf не работает, я не разобрался каво,
		// 	// попросил блядского пидорасного хуесосного чатаджипити помочь, в надежде, что он
		// 	// хоть чуть-чутьь не апездол. но он полнейший апездал и бболее того, если бы он не
		// 	// был апездалом, я бы почти ничего не вынес из обработки png
		// 	// я буквально просрал весь день. ЧАТДЖИПИТИ ПРОСТО ЕБАНЕЙШИЙ ДОЛБОЕБ
		// 	if err := ult.pdf.Work(message.UserId, message.Path); err != nil {
		// 		log.Println("Error operating pdf:", err)
		// 	}
	}
	return nil
}
