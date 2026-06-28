package mp3

import (
	"log"
	"os"
	"worker/domain"
	"worker/interfaces"

	"github.com/hajimehoshi/go-mp3"
)

type AudioUpdate struct {
	Switcher interfaces.Switcher
	Updater  interfaces.MP3Updater
}

func (upd *AudioUpdate) Work(userId int, filepath string) error {

	f, err := os.OpenFile(filepath, os.O_RDONLY, 0644)
	if err != nil {
		upd.Switcher.StatusFail(userId, filepath)
		return err
	}
	defer f.Close()

	log.Println(1)

	decoder, err := mp3.NewDecoder(f)
	if err != nil {
		upd.Switcher.StatusFail(userId, filepath)
		return err
	}

	log.Println(2)

	length := decoder.Length()
	sample := decoder.SampleRate()
	durationInt := int(length) / (sample * 4)
	mp3Data := domain.MP3Data{
		UserId:   userId,
		Filepath: filepath,
		Length:   durationInt,
	}

	log.Println(3)

	if err = upd.Updater.Mp3Udate(mp3Data); err != nil {
		upd.Switcher.StatusFail(userId, filepath)
		return err
	}
	return nil
}
