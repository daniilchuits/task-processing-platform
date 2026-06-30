package repo

import (
	"log"
	"worker/domain"
)

func (repo tasksRepo) TxtUpdate(data domain.DataTxt) error {

	query := `
		UPDATE tasks
		SET 
			phrase_count_txt=$1,
			lines_txt=$2
		WHERE 
			id=$3
				AND filepath=$4
	`

	_, err := repo.db.Exec(
		query,
		data.PhrasesCount,
		data.Lines,
		data.Id,
		data.Filepath,
	)
	if err != nil {
		log.Printf("Updating %s error: %v\n", data.Filepath, err)
		return domain.ErrUpdatingTasks
	}
	return nil
}
