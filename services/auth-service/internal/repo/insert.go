package repo

import "auth/internal/domain"

func (repo repoManager) InsertCredentials(cred domain.Credentials) (int, error) {

	query := `
		INSERT INTO users (login, hashed_password) VALUES
		($1,$2)
		RETURNING id
	`

	var id int
	err := repo.repo.QueryRow(query, cred.Login, cred.Password).Scan(&id)
	return id, err
}
