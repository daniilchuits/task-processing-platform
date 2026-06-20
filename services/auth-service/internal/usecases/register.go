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

func (register RegisterUsecase) Exec(credentials domain.Credentials) (int, error) {

	exists, err := register.UserExists.UserExists(credentials.Login)
	if err != nil {
		log.Println("Error checking if user exists:", err)
		return 0, err
	}
	if exists {
		log.Println(domain.ErrUserExists.Error())
		return 0, domain.ErrUserExists
	}

	// do hashing password epta

	id, err := register.InsertCredentials.InsertCredentials(credentials)
	if err != nil {
		log.Println("Error inserting user:", err)
		return 0, err
	}
	return id, nil
}
