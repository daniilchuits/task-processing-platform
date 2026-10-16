package repo

import "notifycation-service/internal/domain"

func (manager *notifyRepo) UpdateTasks(data domain.Msg) (int, error) {

	query := `
		UPDATE tasks
		SET 
			status=$1,
			error=$2
		WHERE id=$3
	`

	res, err := manager.db.Exec(
		query,
		data.Status,
		data.Err,
		data.Id,
	)
	if err != nil {
		return 0, err
	}
	r, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(r), nil
}
