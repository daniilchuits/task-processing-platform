package domain

import (
	"io"
	"testing"
)

type MockFile struct {
	Err error
}

func (mock MockFile) Close() error {
	return mock.Err
}

func (mock MockFile) Read(p []byte) (n int, err error) {
	if mock.Err == nil {
		return 0, io.EOF
	}
	return 0, mock.Err
}

func (mock MockFile) ReadAt(p []byte, off int64) (n int, err error) {
	return 0, nil
}

func (mock MockFile) Seek(offset int64, whence int) (int64, error) {
	return 0, nil
}

func TestCreateFile(t *testing.T) {

	mock := MockFile{}

	data := []struct {
		name     string
		filename string
		expErr   error
	}{
		{"nil error", "realPath", nil},
		{"empty name so err creating", "", ErrCreating},
	}

	for _, tt := range data {
		t.Run(tt.name, func(t *testing.T) {
			err := CreateFile(tt.filename, mock)
			if err != tt.expErr {
				t.Fatalf("Expected %v err, got %v", tt.expErr, err)
			}
		})
	}
}

func TestCreateFileWErr(t *testing.T) {

	mock := MockFile{Err: ErrSelectingAll}

	data := []struct {
		name     string
		filename string
		expErr   error
	}{
		{"Err selecting all", "realPath", ErrCopy},
		{"Have an err in mock, so Close() will not return EOF, but mock's error", "", ErrCreating},
	}

	for _, tt := range data {
		t.Run(tt.name, func(t *testing.T) {
			err := CreateFile(tt.filename, mock)
			if err != tt.expErr {
				t.Fatalf("Expected %v err, got %v", tt.expErr, err)
			}
		})
	}
}
