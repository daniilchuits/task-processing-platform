package jpg

import (
	"image"
	"image/jpeg"
	"os"
	"strconv"
	"worker/domain"
	"worker/interfaces"
	"worker/rabbitmq"
)

type JpegUpdate struct {
	Switcher interfaces.Switcher
	Update   interfaces.JPGUpdater
	Produce  interfaces.Producer
}

func (upd *JpegUpdate) Work(id int, filepath string) error {

	msg := rabbitmq.Msg{
		Id:     id,
		Status: domain.FailedStatus,
	}

	if err := upd.Switcher.StatusProcessing(id, filepath); err != nil {
		msg.Err = err.Error()
		upd.Produce.Produce(msg)
		return err
	}

	f, err := os.OpenFile(filepath, os.O_RDONLY, 0644)
	if err != nil {
		msg.Err = err.Error()
		upd.Produce.Produce(msg)
		return err
	}
	defer f.Close()

	image, _, err := image.Decode(f)
	if err != nil {
		msg.Err = err.Error()
		upd.Produce.Produce(msg)
		return err
	}

	if _, err = f.Seek(0, 0); err != nil {
		msg.Err = err.Error()
		upd.Produce.Produce(msg)
		return err
	}

	config, err := jpeg.DecodeConfig(f)
	if err != nil {
		msg.Err = err.Error()
		upd.Produce.Produce(msg)
		return err
	}

	resolution := strconv.Itoa(config.Height) + "x" + strconv.Itoa(config.Width)

	mainColors := process(image)

	jpgData := domain.DataJPG{
		Id:         id,
		Filepath:   filepath,
		Resolution: resolution,
		MainColors: mainColors,
	}

	if err = upd.Update.JPGUpdate(jpgData); err != nil {
		msg.Err = err.Error()
		upd.Produce.Produce(msg)
		return err
	}

	msg.Status = domain.FinishedStatus
	upd.Produce.Produce(msg)

	return nil
}
