package jpg

import (
	"image"
	"image/jpeg"
	"os"
	"strconv"
	"worker/domain"
	"worker/interfaces"
)

type JpegUpdate struct {
	Switcher interfaces.Switcher
	Update   interfaces.JPGUpdater
}

func (upd *JpegUpdate) Work(userId int, filepath string) error {

	if err := upd.Switcher.StatusProcessing(userId, filepath); err != nil {
		upd.Switcher.StatusFail(userId, filepath)
		return err
	}

	f, err := os.OpenFile(filepath, os.O_RDONLY, 0644)
	if err != nil {
		upd.Switcher.StatusFail(userId, filepath)
		return err
	}
	defer f.Close()

	image, _, err := image.Decode(f)
	if err != nil {
		upd.Switcher.StatusFail(userId, filepath)
		return err
	}

	if _, err = f.Seek(0, 0); err != nil {
		return err
	}

	config, err := jpeg.DecodeConfig(f)
	if err != nil {
		upd.Switcher.StatusFail(userId, filepath)
		return err
	}

	resolution := strconv.Itoa(config.Height) + "x" + strconv.Itoa(config.Width)

	mainColors := process(image)

	jpgData := domain.DataJPG{
		UserId:     userId,
		Filepath:   filepath,
		Resolution: resolution,
		MainColors: mainColors,
	}

	if err = upd.Update.JPGUpdate(jpgData); err != nil {
		upd.Switcher.StatusFail(userId, filepath)
		return err
	}

	return nil
}
