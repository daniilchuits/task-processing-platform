package domain

import (
	"regexp"
)

const (
	minLenLogin    = 3
	minLenPassword = 8
	maxLen         = 30
)

var loginRegexp = regexp.MustCompile("^[a-zA-Z0-9_-]+$")

func Validating(cred Credentials) error {

	if len(cred.Login) < minLenLogin || len(cred.Login) > maxLen {
		return ErrLogin
	}

	if len(cred.Password) < minLenPassword || len(cred.Password) > maxLen {
		return ErrPassword
	}

	if !loginRegexp.MatchString(cred.Login) {
		return ErrForbiddenSymbols
	}
	return nil
}
