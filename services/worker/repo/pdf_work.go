package repo

import (
	"log"
	"worker/domain"
)

func (repo tasksRepo) PdfUpdate(data domain.PDFData) error {

	query := `
		UPDATE tasks
		SET
			pages_pdf=$1,
			photo_in_doc_pdf=$2
		WHERE user_id=$3
			AND filepath=$4
	`

	res, err := repo.db.Exec(
		query,
		data.Pages,
		data.HasImage,
		data.UserId,
		data.Filepath,
	)
	if err != nil {
		log.Printf("Updating %s error: %v\n", data.Filepath, err)
		return domain.ErrUpdatingTasks
	}
	rows, _ := res.RowsAffected()
	log.Println("Rows affected:", rows)
	log.Println(7)
	return err
}
