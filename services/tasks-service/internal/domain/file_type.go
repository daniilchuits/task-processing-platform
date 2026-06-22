package domain

import (
	"log"
	"strings"
)

func DetermineType(filename string) (string, error) {

	count := strings.Count(filename, ".")
	if count != 1 {
		log.Println("'.' count=", count)
		return "", ErrInvalidExtension
	}

	switch {

	case strings.HasSuffix(filename, jpgExtension):
		return jpg, nil
	case strings.HasSuffix(filename, zipExtension):
		return zip, nil
	case strings.HasSuffix(filename, txtExtension):
		return txt, nil
	case strings.HasSuffix(filename, mp3Extension):
		return mp3, nil
	case strings.HasSuffix(filename, csvExtension):
		return csv, nil
	case strings.HasSuffix(filename, pdfExtension):
		return pdf, nil
	default:
		return "", ErrInvalidExtension
	}
}
