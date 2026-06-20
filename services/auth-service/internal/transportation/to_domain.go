package transportation

import "auth/internal/domain"

func ToDomain(credHttp Credentials) domain.Credentials {
	return domain.Credentials{
		Id:       credHttp.Id,
		Login:    credHttp.Login,
		Password: credHttp.Password,
	}
}
