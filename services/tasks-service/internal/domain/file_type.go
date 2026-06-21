package domain

import "strings"

func DetermineType(filename string) (string, error) {

	count := strings.Count(filename, ".")
	if count != 1 {
		return "", ErrInvalidExtension
	}

	switch {

	case strings.HasPrefix(filename, jpgExtension):
		return jpg, nil
	case strings.HasPrefix(filename, zipExtension):
		return zip, nil
	case strings.HasPrefix(filename, txtExtension):
		return txt, nil
	case strings.HasPrefix(filename, mp3Extension):
		return mp3, nil
	case strings.HasPrefix(filename, csvExtension):
		return csv, nil
	case strings.HasPrefix(filename, pdfExtension):
		return pdf, nil
	default:
		return "", ErrInvalidExtension
	}
}
