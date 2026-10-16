package repo

import (
	"auth/internal/domain"
	"testing"
)

func (repo repoManagerMock) InsertCredentials(credentials domain.Credentials) (*domain.Credentials, error) {
	return &credentials, nil
}

func (repo repoManagerMockWErrUserExists) InsertCredentials(credentials domain.Credentials) (*domain.Credentials, error) {
	return nil, repo.err
}

func TestInsertCredentials(t *testing.T) {

	data := []struct {
		name   string
		cred   domain.Credentials
		expErr error
	}{
		{name: "norm", cred: domain.Credentials{Login: "danilkazx", Password: "normal-password"}, expErr: nil},
		{name: "empty creds will not be here, so it will be nil err", cred: domain.Credentials{}, expErr: nil},
		{name: "toze normal, but empty password", cred: domain.Credentials{Login: "dan228play", Password: ""}, expErr: nil},
	}

	mock := repoManagerMock{}

	for _, tt := range data {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := mock.InsertCredentials(tt.cred); err != tt.expErr {
				t.Errorf("Wanted err %v, got %v\n", tt.expErr, err)
			}
		})
	}
}

func TestInsertCredentialsWErr(t *testing.T) {

	data := []struct {
		name   string
		cred   domain.Credentials
		expErr error
	}{
		{name: "Creds aren't important mock's error", cred: domain.Credentials{Login: "danilkazx", Password: "normal-password"}, expErr: domain.RepoMockErr},
		{name: "Zero-value creds, mock's error", cred: domain.Credentials{}, expErr: domain.RepoMockErr},
	}

	mock := repoManagerMockWErrUserExists{err: domain.RepoMockErr}

	for _, tt := range data {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := mock.InsertCredentials(tt.cred); err != tt.expErr {
				t.Errorf("Expected %v err, got %v", tt.expErr, err)
			}
		})
	}
}
