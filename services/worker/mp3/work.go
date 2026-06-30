package mp3

import (
	"os"
	"worker/domain"
	"worker/interfaces"
	"worker/rabbitmq"

	"github.com/hajimehoshi/go-mp3"
)

type AudioUpdate struct {
	Switcher interfaces.Switcher
	Updater  interfaces.MP3Updater
	Producer interfaces.Producer
}

func (upd *AudioUpdate) Work(id int, filepath string) error {

	msg := rabbitmq.Msg{
		Id:     id,
		Status: domain.FailedStatus,
	}

	if err := upd.Switcher.StatusProcessing(id, filepath); err != nil {
		msg.Err = err.Error()
		upd.Producer.Produce(msg)
		return err
	}

	f, err := os.OpenFile(filepath, os.O_RDONLY, 0644)
	if err != nil {
		msg.Err = err.Error()
		upd.Producer.Produce(msg)
		return err
	}
	defer f.Close()

	decoder, err := mp3.NewDecoder(f)
	if err != nil {
		msg.Err = err.Error()
		upd.Producer.Produce(msg)
		return err
	}

	length := decoder.Length()
	sample := decoder.SampleRate()
	durationInt := int(length) / (sample * 4)
	mp3Data := domain.MP3Data{
		Id:       id,
		Filepath: filepath,
		Length:   durationInt,
	}

	if err = upd.Updater.Mp3Udate(mp3Data); err != nil {
		msg.Err = err.Error()
		upd.Producer.Produce(msg)
		return err
	}

	msg.Status = domain.FinishedStatus
	upd.Producer.Produce(msg)
	return nil
}
