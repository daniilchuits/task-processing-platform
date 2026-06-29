package csv

import (
	"encoding/csv"
	"os"
	"worker/domain"
	"worker/interfaces"
)

type CSVUpdater struct {
	Switcher interfaces.Switcher
	Update   interfaces.CSVUpdater
}

func (upd *CSVUpdater) Work(userId int, filepath string) error {

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

	csvDecoder := csv.NewReader(f)
	records, err := csvDecoder.ReadAll()
	if err != nil {
		upd.Switcher.StatusFail(userId, filepath)
		return err
	}

	numLines := len(records)

	data := domain.CsvData{
		UserId:   userId,
		Filepath: filepath,
		Lines:    numLines,
	}

	if err = upd.Update.CSVUpdate(data); err != nil {
		upd.Switcher.StatusFail(userId, filepath)
		return err
	}
	return nil
}
