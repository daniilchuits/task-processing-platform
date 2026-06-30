package repo

import (
	"log"
	"worker/domain"
)

func (repo tasksRepo) StatusProcessing(id int, path string) error {

	query := `
		UPDATE tasks
		SET status=$1
		WHERE id=$2
			AND filepath=$3
	`

	if _, err := repo.db.Exec(
		query,
		domain.ProcessingStatus,
		id,
		path,
	); err != nil {
		log.Printf("Updating status to 'processing' file %s error: %v\n", path, err)
		return domain.ErrProcesing
	}
	return nil
}
