package domain

import "golang.org/x/crypto/bcrypt"

const hashCost = 10

func Hashing(password string) ([]byte, error) {

	return bcrypt.GenerateFromPassword([]byte(password), hashCost)
}

func CompareHashedPassword(tablePassword, writtenPassword string) error {

	return bcrypt.CompareHashAndPassword([]byte(tablePassword), []byte(writtenPassword))
}
