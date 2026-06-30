package repo

import (
	"log"
	"worker/domain"
)

func (repo tasksRepo) ZipUpdate(data domain.ZipData) error {

	query := `
		UPDATE tasks
		SET 
			size_after_unzip_zip=$1,
			zip_files_zip=$2
		WHERE id=$3
			AND filepath=$4
	`

	_, err := repo.db.Exec(
		query,
		data.Size,
		data.Files,
		data.Id,
		data.Filepath,
	)
	if err != nil {
		log.Printf("Updating %s error: %v\n", data.Filepath, err)
		return domain.ErrUpdatingTasks
	}
	return nil
}
