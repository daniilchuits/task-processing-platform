package repo

func (manager *repoManager) Check(userID int, filename string) (bool, error) {

	query := `
		SELECT EXISTS(
			SELECT 1
			FROM tasks
			WHERE user_id=$1
				AND filename=$2
		)
	`

	var exists bool
	err := manager.db.QueryRow(query, userID, filename).Scan(
		&exists,
	)
	return exists, err
}
