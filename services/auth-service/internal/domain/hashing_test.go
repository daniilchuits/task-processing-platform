package domain

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestHashing(t *testing.T) {
	data := []struct {
		name     string
		password string
		expErr   error
	}{
		{
			name:     "no err",
			password: "QWELUMT*",
			expErr:   nil,
		},
		{
			name:     "empty password",
			password: "",
			expErr:   nil,
		},
	}

	for _, tt := range data {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Hashing(tt.password)
			if err != tt.expErr {
				t.Errorf("Want %v, got %v\n", tt.expErr, err)
			}
		})
	}
}

func TestCompareHashedPassword(t *testing.T) {

	samePsStr := "zxcasdqwe"
	samePs, err := bcrypt.GenerateFromPassword([]byte(samePsStr), hashCost)
	if err != nil {
		t.Fatalf("Hashing password error: %v\n", err)
	}

	diffPsStr := "qweasdzxcc"
	diffPs, err := bcrypt.GenerateFromPassword([]byte(diffPsStr), hashCost)
	if err != nil {
		t.Fatalf("Hashing password error: %v\n", err)
	}

	data := []struct {
		name     string
		tableP   string
		writtenP string
		err      error
	}{
		{
			"same Ps", string(samePs), "zxcasdqwe", nil,
		},
		{
			"diff Ps", string(diffPs), "fefafafqwe", bcrypt.ErrMismatchedHashAndPassword,
		},
	}

	for _, tt := range data {
		t.Run(tt.name, func(t *testing.T) {
			if err := CompareHashedPassword(tt.tableP, tt.writtenP); err != tt.err {
				t.Errorf("Wanted %v, got %v\n", tt.err, err)
			}
		})
	}
}
