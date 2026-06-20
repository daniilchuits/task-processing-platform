package transportation

import "auth/internal/domain"

func ToHttp(id int, cred domain.Credentials) Credentials {
	return Credentials{
		Id:       id,
		Login:    cred.Login,
		Password: cred.Password,
	}
}
