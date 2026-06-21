package repo

func (repo *repoManager) SelectPassword(login string) (string, int, error) {

	query := `
		SELECT 
			hashed_password,
			id
		FROM users
		WHERE login=$1
	`

	var (
		password string
		id       int
	)
	err := repo.repo.QueryRow(query, login).Scan(&password, &id)
	return password, id, err
}
