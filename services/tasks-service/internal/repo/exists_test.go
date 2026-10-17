package repo

import "testing"

func TestCheck(t *testing.T) {

	mock := repoMock{}
	repoManager := NewRepoManager(mock)

	t.Run("ебать мок", func(t *testing.T) {
		exists, err := repoManager.Check(5, "z")
		if err != nil {
			t.Fatalf("Expected nil err, got %v", err)
		}
		if !exists {
			t.Errorf("Expected exists == true, got %t", exists)
		}
	})
}
