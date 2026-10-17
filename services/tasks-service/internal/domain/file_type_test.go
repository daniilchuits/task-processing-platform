package domain

import "testing"

func TestDetermineType(t *testing.T) {

	data := []struct {
		name     string
		filename string
		filetype string
		expErr   error
	}{
		{
			"empty name", "", "", ErrInvalidExtension,
		},
		{
			"2 dots", "file..txt", "", ErrInvalidExtension,
		},
		{
			"only dots", ".......", "", ErrInvalidExtension,
		},
		{
			"jpg", "file.jpg", jpg, nil,
		},
		{
			"mp3", "file.mp3", mp3, nil,
		},
		{
			"csv", "file.csv", csv, nil,
		},
		{
			"pdf", "file.pdf", pdf, nil,
		},
		{
			"txt", "file.txt", txt, nil,
		},
		{
			"zip", "file.zip", zip, nil,
		},
	}

	for _, tt := range data {
		t.Run(tt.name, func(t *testing.T) {
			res, err := DetermineType(tt.filename)
			if err != tt.expErr {
				t.Fatalf("Expected %v err, got %v", tt.expErr, err)
			}
			if res != tt.filetype {
				t.Errorf("Expected type %s, got %s", tt.filetype, res)
			}
		})
	}
}
