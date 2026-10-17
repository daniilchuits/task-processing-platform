package domain

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

const UploadDir = "/app/uploads"

func CreateFile(filename string, file multipart.File) error {
	defer file.Close()

	fPath := filepath.Join(".", UploadDir)
	fPathCurr := filepath.Join(fPath, filename)

	if err := os.MkdirAll(fPath, 0755); err != nil {
		return err
	}
	f, err := os.OpenFile(fPathCurr, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return ErrCreating
	}
	defer f.Close()

	// filepathCurr := filepath.Join(UploadDir, filename)

	// if _, err := os.OpenFile(filepathCurr, os.O_RDONLY, 0600); err == nil {
	// 	return ErrFileExists
	// }

	// fPath := filepath.Join(".", UploadDir)
	// err := os.MkdirAll(fPath, 0755)
	// if err != nil {
	// 	return err
	// }

	// f, err := os.OpenFile(filepathCurr, os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	return ErrCreating
	// }
	// defer f.Close()

	if _, err = io.Copy(f, file); err != nil {
		return ErrCopy
	}
	return nil
}
