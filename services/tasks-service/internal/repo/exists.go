package repo

func (repo *repoManager) NoteExists(userId int, filename string) (bool, error) {

	query := `
		SELECT EXISTS(
			SELECT 1
			FROM tasks
			WHERE user_id=$1
				AND filename=$2
		);
	`

	var exists bool
	err := repo.db.QueryRow(query, userId, filename).Scan(&exists)
	return exists, err
}
