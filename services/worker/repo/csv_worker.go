package repo

import (
	"log"
	"worker/domain"
)

func (repo tasksRepo) CSVUpdate(data domain.CsvData) error {

	query := `
		UPDATE tasks
		SET num_of_lines_csv=$1
		WHERE id=$2
			AND filepath=$3
	`

	_, err := repo.db.Exec(
		query,
		data.Lines,
		data.Id,
		data.Filepath,
	)
	if err != nil {
		log.Printf("Updating %s error: %v\n", data.Filepath, err)
		return domain.ErrUpdatingTasks
	}
	return err
}
