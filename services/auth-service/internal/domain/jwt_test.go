package domain

import "testing"

func TestMakeJWT(t *testing.T) {

	secret := NewSecretKey("secret")
	data := []struct {
		name   string
		id     int
		expErr error
	}{
		{
			"ok", 1, nil,
		},
		{
			"zero id", 0, nil,
		},
		{
			"huge id", 100000000000000000, nil,
		},
	}

	for _, tt := range data {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := secret.MakeJWT(tt.id); err != tt.expErr {
				t.Errorf("Waited %v, got %v\n", tt.expErr, err)
			}
		})
	}
}
