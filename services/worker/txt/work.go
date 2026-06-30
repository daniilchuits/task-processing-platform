package txt

import (
	"io"
	"os"
	"strings"
	"worker/domain"
	"worker/interfaces"
	"worker/rabbitmq"
)

type TxtUpdate struct {
	Txt      interfaces.TxtUpdater
	Switcher interfaces.Switcher
	Produce  interfaces.Producer
}

func (txt *TxtUpdate) Work(id int, filepath string) error {

	msg := rabbitmq.Msg{
		Id:     id,
		Status: domain.FailedStatus,
	}

	if err := txt.Switcher.StatusProcessing(id, filepath); err != nil {
		msg.Err = err.Error()
		txt.Produce.Produce(msg)
		return err
	}

	f, err := os.OpenFile(filepath, os.O_RDONLY, 0644)
	if err != nil {
		msg.Err = err.Error()
		txt.Produce.Produce(msg)
		return err
	}
	defer f.Close()

	byter, err := io.ReadAll(f)
	if err != nil {
		msg.Err = err.Error()
		txt.Produce.Produce(msg)
		return err
	}

	strs := string(byter)

	lines := strings.Count(strs, "\n") + 1
	words := len(strings.Fields(strs))

	data := domain.DataTxt{
		Id:           id,
		Filepath:     filepath,
		Lines:        lines,
		PhrasesCount: words,
	}

	if err = txt.Txt.TxtUpdate(data); err != nil {
		msg.Err = err.Error()
		txt.Produce.Produce(msg)
		return domain.ErrUpdatingTasks
	}

	msg.Status = domain.FinishedStatus
	txt.Produce.Produce(msg)

	return nil
}
