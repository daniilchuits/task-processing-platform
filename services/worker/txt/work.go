package txt

import (
	"io"
	"log"
	"os"
	"strings"
	"worker/domain"
	"worker/interfaces"
)

type TxtUpdate struct {
	Txt      interfaces.TxtUpdater
	Switcher interfaces.Switcher
}

func (txt *TxtUpdate) Work(userId int, filepath string) error {

	if err := txt.Switcher.StatusProcessing(userId, filepath); err != nil {
		return err
	}

	f, err := os.OpenFile(filepath, os.O_RDONLY, 0644)
	if err != nil {
		txt.Switcher.StatusFail(userId, filepath)
		return err
	}
	defer f.Close()

	byter, err := io.ReadAll(f)
	if err != nil {
		txt.Switcher.StatusFail(userId, filepath)
		return err
	}

	strs := string(byter)

	lines := strings.Count(strs, "\n") + 1
	words := len(strings.Fields(strs))

	data := domain.DataTxt{
		UserId:       userId,
		Filepath:     filepath,
		Lines:        lines,
		PhrasesCount: words,
	}

	if err = txt.Txt.TxtUpdate(data); err != nil {
		log.Println("Updating tasks txt err:", err)
		return domain.ErrUpdatingTasks
	}
	return nil
}
