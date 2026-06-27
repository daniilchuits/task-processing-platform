package repo

import (
	"time"
	"worker/domain"
)

func (repo tasksRepo) TxtUpdate(data domain.DataTxt) error {

	query := `
		UPDATE tasks
		SET 
			phrase_count_txt=$1,
			lines_txt=$2
		WHERE 
			user_id=$3
				AND filepath=$4
	`

	_, err := repo.db.Exec(query, data.PhrasesCount, data.Lines, data.UserId, data.Filepath)
	time.Sleep(time.Second * 2)
	return err
}
