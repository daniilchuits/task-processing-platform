package repo

import "auth/internal/domain"

func (repo repoManager) InsertCredentials(credentials domain.Credentials) (*domain.Credentials, error) {

	query := `
		INSERT INTO users (login, hashed_password) VALUES
		($1,$2)
		RETURNING id, login, hashed_password
	`

	var cred domain.Credentials
	err := repo.repo.QueryRow(query, credentials.Login, credentials.Password).Scan(
		&cred.Id,
		&cred.Login,
		&cred.Password,
	)
	return &cred, err
}
