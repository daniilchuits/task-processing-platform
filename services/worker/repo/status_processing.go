package repo

import "worker/domain"

func (repo tasksRepo) StatusProcessing(userId int, path string) error {

	query := `
		UPDATE tasks
		SET status=$1
		WHERE user_id=$2
			AND filepath=$3
	`

	if _, err := repo.db.Exec(
		query,
		domain.ProcessingStatus,
		userId,
		path,
	); err != nil {
		return domain.ErrProcesing
	}
	return nil
}
