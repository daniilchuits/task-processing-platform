package repo

import "worker/domain"

func (repo tasksRepo) StatusFail(userId int, path string) error {

	query := `
		UPDATE tasks
		SET status=$1
		WHERE user_id=$2
			AND filepath=$3
		RETURNING 
			id,
			user_id,
			filename,
			filepath,
			status,
			type
	`

	if _, err := repo.db.Exec(
		query,
		domain.FailedStatus,
		userId,
		path,
	); err != nil {
		return domain.ErrFailed
	}
	return nil
}
