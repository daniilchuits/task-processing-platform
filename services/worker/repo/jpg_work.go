package repo

import (
	"log"
	"worker/domain"
)

func (repo tasksRepo) JPGUpdate(data domain.DataJPG) error {

	query := `
		UPDATE tasks
		SET
			main_colors_jpg=$1,
			photo_resolution_jpg=$2
		WHERE user_id=$3
			AND filepath=$4
	`

	_, err := repo.db.Exec(
		query,
		data.MainColors,
		data.Resolution,
		data.UserId,
		data.Filepath,
	)
	if err != nil {
		log.Printf("Updating %s error: %v\n", data.Filepath, err)
		return domain.ErrUpdatingTasks
	}
	return err
}
