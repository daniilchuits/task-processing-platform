package repo

import "notifycation-service/internal/domain"

func (manager *notifyRepo) UpdateTasks(data domain.Msg) error {

	query := `
		UPDATE tasks
		SET 
			status=$1,
			error=$2
		WHERE id=$3
	`

	_, err := manager.db.Exec(
		query,
		data.Status,
		data.Err,
		data.Id,
	)
	return err
}
