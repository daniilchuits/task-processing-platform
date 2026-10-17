package repo

import (
	"errors"
	"testing"
)

func TestInsert(t *testing.T) {

	mock := repoMock{}
	dbManager := NewRepoManager(mock)

	data := []struct {
		name     string
		userID   int
		filename string
		filepath string
		typ      string
		expErr   error
	}{
		{"no err", 1, "file.txt", "filepath.txt", "txt", nil},
		{"no err again", 2, "file.mp3", "filepath.mp3", "mp3", nil},
	}

	for _, tt := range data {
		t.Run(tt.name, func(t *testing.T) {
			_, err := dbManager.Insert(tt.userID, tt.filename, tt.filepath, tt.typ)
			if err != tt.expErr {
				t.Fatalf("Expected %v, got %v", tt.expErr, err)
			}
		})
	}
}

func TestInsertWErr(t *testing.T) {

	mock := repoMock{err: PostingErr}
	dbManager := NewRepoManager(mock)

	data := []struct {
		name     string
		userID   int
		filename string
		filepath string
		typ      string
		expErr   error
	}{
		{"PostingErr in mock", 1, "file.txt", "filepath.txt", "txt", PostingErr},
		{"PostingErr in mock again", 0, "", "", "", PostingErr},
	}

	for _, tt := range data {
		t.Run(tt.name, func(t *testing.T) {
			_, err := dbManager.Insert(tt.userID, tt.filename, tt.filepath, tt.typ)
			if !errors.Is(err, PostingErr) {
				t.Fatalf("Expected %v, got %v", tt.expErr, err)
			}
		})
	}
}
