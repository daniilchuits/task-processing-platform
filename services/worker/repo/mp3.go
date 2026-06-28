package repo

import (
	"log"
	"worker/domain"
)

func (repo tasksRepo) Mp3Udate(data domain.MP3Data) error {

	query := `
		UPDATE tasks
		SET audio_length_mp3=$1
		WHERE user_id=$2
			AND filepath=$3
	`

	_, err := repo.db.Exec(
		query,
		data.Length,
		data.UserId,
		data.Filepath,
	)
	if err != nil {
		log.Printf("Updating %s error: %v\n", data.Filepath, err)
		return domain.ErrUpdatingTasks
	}
	return err
}
