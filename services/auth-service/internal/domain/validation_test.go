package domain

import "testing"

func TestValidation(t *testing.T) {

	data := []struct {
		name string
		Credentials
		err error
	}{
		{
			name: "empty login",
			Credentials: Credentials{
				Login:    "",
				Password: "",
			},
			err: ErrLogin,
		},
		{
			name: "long login",
			Credentials: Credentials{
				Login:    "qweqweqweqweqweqweqweqweqweqwe",
				Password: "",
			},
			err: ErrLogin,
		},
		{
			name: "short password",
			Credentials: Credentials{
				Login:    "danilka228",
				Password: "zxcasd",
			},
			err: ErrPassword,
		},
		{
			name: "Long password",
			Credentials: Credentials{
				Login:    "danilka228",
				Password: "qweqweqweqweqweqweqweqweqweqwe",
			},
			err: ErrPassword,
		},
		{
			name: "forbidden symbs",
			Credentials: Credentials{
				Login:    "danilka228@play.com",
				Password: "qwer12340659",
			},
			err: ErrForbiddenSymbols,
		},
		{
			name: "no error",
			Credentials: Credentials{
				Login:    "danik228play",
				Password: "Lmtdfmlygy228",
			},
			err: nil,
		},
	}

	for _, tt := range data {
		t.Run(tt.name, func(t *testing.T) {
			err := Validating(tt.Credentials)
			if err != tt.err {
				t.Errorf("Wanted %v, got %v\n", tt.err, err)
			}
		})
	}
}
