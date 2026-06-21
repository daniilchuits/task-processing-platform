package transportation

import "auth/internal/domain"

func ToHttp(cred domain.Credentials) Credentials {
	return Credentials{
		Id:       cred.Id,
		Login:    cred.Login,
		Password: cred.Password,
	}
}
