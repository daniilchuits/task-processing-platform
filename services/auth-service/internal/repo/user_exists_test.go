package repo

import "testing"

func (repo repoManagerMock) UserExists(login string) (bool, error) {
	return false, nil
}

func (repo repoManagerMockWErrUserExists) UserExists(login string) (bool, error) {
	return repo.exists, nil
}

func TestUserExists(t *testing.T) {

	mock := repoManagerMock{}

	data := []struct {
		name      string
		login     string
		expExists bool
		expErr    error
	}{
		{"normal login not exists", "danilkazxc", false, nil},
		{"empty login not exists", "", false, nil},
	}

	for _, tt := range data {
		t.Run(tt.name, func(t *testing.T) {
			exists, err := mock.UserExists(tt.login)
			if err != tt.expErr {
				t.Fatalf("Expected %v err, got %v", tt.expErr, err)
			}
			if exists != tt.expExists {
				t.Errorf("Expected existence %t, got %t", tt.expExists, exists)
			}
		})
	}
}

func TestUserExistsExists(t *testing.T) {

	mock := repoManagerMockWErrUserExists{
		exists: true,
	}

	data := []struct {
		name      string
		login     string
		expExists bool
		expErr    error
	}{
		{"normal login not exists", "danilkazxc", true, nil},
		{"empty login not exists", "", true, nil},
	}

	for _, tt := range data {
		t.Run(tt.name, func(t *testing.T) {
			exists, err := mock.UserExists(tt.login)
			if err != tt.expErr {
				t.Fatalf("Expected %v err, got %v", tt.expErr, err)
			}
			if exists != tt.expExists {
				t.Errorf("Expected existence %t, got %t", tt.expExists, exists)
			}
		})
	}
}
