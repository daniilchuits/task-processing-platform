package usecases

import (
	"auth/internal/domain"
	"auth/internal/interfaces"
	"log"
)

type RegisterUsecase struct {
	UserExists        interfaces.UserExistsInterface
	InsertCredentials interfaces.InsertCredentialsInterface
}

func (register RegisterUsecase) Exec(credentials domain.Credentials) (*domain.Credentials, error) {

	exists, err := register.UserExists.UserExists(credentials.Login)
	if err != nil {
		log.Println("Error checking if user exists:", err)
		return nil, err
	}
	if exists {
		log.Println(domain.ErrUserExists.Error())
		return nil, domain.ErrUserExists
	}

	hashedPassword, err := domain.Hashing(credentials.Password)
	if err != nil {
		log.Println("Hashing err:", err)
		return nil, domain.ErrDuringHashing
	}
	credentials.Password = string(hashedPassword)

	cred, err := register.InsertCredentials.InsertCredentials(credentials)
	if err != nil {
		log.Println("Error inserting user:", err)
		return nil, err
	}
	return cred, nil
}
