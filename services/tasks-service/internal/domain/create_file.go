package domain

import (
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
)

const UploadDir = "/app/uploads"

func CreateFile(filename string, file multipart.File) error {
	defer file.Close()

	filepath := filepath.Join(UploadDir, filename)

	if _, err := os.OpenFile(filepath, 0600, os.FileMode(os.O_RDONLY)); err == nil {
		log.Println("Error checking file existence:", err)
		return ErrFileExists
	}

	f, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Error creating file:", err)
		return ErrCreating
	}
	defer f.Close()

	if _, err = io.Copy(f, file); err != nil {
		log.Println("Error to copy file:", err)
		return ErrCopy
	}
	return nil
}
