package repo

import (
	"auth/internal/domain"
	"testing"
)

func (repo repoManagerMock) SelectPassword(login string) (string, int, error) {
	return "", 0, nil
}

func (repo repoManagerMockWErrUserExists) SelectPassword(login string) (string, int, error) {
	return "", 0, repo.err
}

func TestSelectPassword(t *testing.T) {
	mock := repoManagerMock{}

	data := []struct {
		name   string
		login  string
		expErr error
	}{
		{
			"empty login nil error", "", nil,
		},
		{
			"normal login nil error", "danilkazxc", nil,
		},
	}

	for _, tt := range data {
		t.Run(tt.name, func(t *testing.T) {
			if _, _, err := mock.SelectPassword(tt.login); err != tt.expErr {
				t.Errorf("Expected %v err, got %v", tt.expErr, err)
			}
		})
	}
}

func TestSelectPasswordWMockErr(t *testing.T) {
	mock := repoManagerMockWErrUserExists{err: domain.RepoMockErr}

	data := []struct {
		name   string
		login  string
		expErr error
	}{
		{
			"empty login mock err", "", domain.RepoMockErr,
		},
		{
			"normal login mock error", "danilkazxc", domain.RepoMockErr,
		},
	}

	for _, tt := range data {
		t.Run(tt.name, func(t *testing.T) {
			if _, _, err := mock.SelectPassword(tt.login); err != tt.expErr {
				t.Errorf("Expected %v err, got %v", tt.expErr, err)
			}
		})
	}
}
