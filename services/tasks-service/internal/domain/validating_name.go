package domain

import "regexp"

var forFilename regexp.Regexp = *regexp.MustCompile("^[a-zA-Z0-9.]+$")

func FilenameValidation(file Filename) error {
	if ok := forFilename.Match([]byte(file.Name)); !ok {
		return ErrRegexp
	}
	return nil
}
