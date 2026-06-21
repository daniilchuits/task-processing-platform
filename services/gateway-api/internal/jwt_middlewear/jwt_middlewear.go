package jwtmiddlewear

import (
	"gateway/internal/domain"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type secretKey struct {
	secret []byte
}

func NewSecret(jwt string) secretKey {
	return secretKey{secret: []byte(jwt)}
}

func (scr secretKey) JwtMiddlewear(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authorization := r.Header.Get("Authorization")

		tokenStr, ok := strings.CutPrefix(authorization, "Bearer ")
		if !ok {
			http.Error(w, domain.ErrNoBearerPrefix.Error(), 400)
			return
		}

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, domain.ErrWrongTokenMethod
			}
			return scr.secret, nil
		})
		if err != nil {
			http.Error(w, domain.ErrWrongTokenMethod.Error(), 400)
			return
		}
		if !token.Valid {
			http.Error(w, domain.ErrTokenInvalid.Error(), 400)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, domain.ErrGettingClaimsFromJWT.Error(), 400)
			return
		}

		userId, ok := claims["user_id"].(float64)
		if !ok {
			http.Error(w, domain.ErrNotFound.Error(), 404)
			return
		}

		r.Header.Del("user_id")
		r.Header.Set("user_id", strconv.Itoa(int(userId)))
	})
}
