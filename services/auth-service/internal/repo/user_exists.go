package repo

func (repo repoManager) UserExists(login string) (bool, error) {

	query := `
		SELECT EXISTS(
			SELECT 1
			FROM users
			WHERE login=$1
		)
	`

	var exists bool
	err := repo.repo.QueryRow(query, login).Scan(&exists)
	return exists, err
}
