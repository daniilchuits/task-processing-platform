package domain

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type secretKey struct {
	key []byte
}

func NewSecretKey(secret string) secretKey {
	return secretKey{key: []byte(secret)}
}

func (scr secretKey) MakeJWT(userId int) (string, error) {

	user_id := strconv.Itoa(userId)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(2 * time.Hour).Unix(),
		"user_id": user_id,
	})
	return token.SignedString(scr.key)
}
