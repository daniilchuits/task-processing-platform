package usecases

import (
	"auth/internal/domain"
	"auth/internal/interfaces"
	"log"
)

type LoginUsecase struct {
	UserExists     interfaces.UserExistsInterface
	SelectPassword interfaces.SelPasswordInterface
	Secret         string
}

func (login LoginUsecase) Exec(cred domain.Credentials) (string, error) {

	exists, err := login.UserExists.UserExists(cred.Login)
	if err != nil {
		log.Println("Err checking existence:", err)
		return "", err
	}
	if !exists {
		return "", domain.ErrUserNotExists
	}

	tablePassword, userId, err := login.SelectPassword.SelectPassword(cred.Login)
	if err != nil {
		log.Println("Error getting table password:", err)
		return "", err
	}

	if err = domain.CompareHashedPassword(tablePassword, cred.Password); err != nil {
		return "", domain.ErrWrongPassword
	}

	newSecret := domain.NewSecretKey(login.Secret)

	jwt, err := newSecret.MakeJWT(userId)
	if err != nil {
		log.Println("Error making jwt:", err)
		return "", domain.ErrCreatingJWT
	}

	return jwt, nil
}
