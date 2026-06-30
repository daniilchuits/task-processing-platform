package zipfile

import (
	"archive/zip"
	"os"
	"worker/domain"
	"worker/interfaces"
	"worker/rabbitmq"
)

type ZipUpdate struct {
	Switcher interfaces.Switcher
	Updater  interfaces.ZipUpdater
	Producer interfaces.Producer
}

func (upd *ZipUpdate) Work(id int, filepath string) error {

	msg := rabbitmq.Msg{
		Id:     id,
		Status: domain.FailedStatus,
	}

	if err := upd.Switcher.StatusProcessing(id, filepath); err != nil {
		msg.Err = err.Error()
		upd.Producer.Produce(msg)
		return err
	}

	f, err := os.Open(filepath)
	if err != nil {
		msg.Err = err.Error()
		upd.Producer.Produce(msg)
		return err
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		msg.Err = err.Error()
		upd.Producer.Produce(msg)
		return err
	}

	reader, err := zip.NewReader(f, info.Size())
	if err != nil {
		msg.Err = err.Error()
		upd.Producer.Produce(msg)
		return err
	}

	files, size := process(reader)
	data := domain.ZipData{
		Id:       id,
		Filepath: filepath,
		Size:     size,
		Files:    files,
	}

	if err = upd.Updater.ZipUpdate(data); err != nil {
		msg.Err = err.Error()
		upd.Producer.Produce(msg)
		return err
	}

	msg.Status = domain.FinishedStatus
	upd.Producer.Produce(msg)
	return nil
}
