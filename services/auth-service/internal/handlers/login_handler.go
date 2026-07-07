package handlers

import (
	"auth/internal/domain"
	"auth/internal/interfaces"
	"auth/internal/transportation"
	"auth/internal/usecases"
	"encoding/json"
	"errors"
	"net/http"
)

type loginHandler struct {
	uc     usecases.LoginUsecase
	secret string
}

func NewLoginHandler(
	exists interfaces.UserExistsInterface,
	sel interfaces.SelPasswordInterface,
	sec string,
) *loginHandler {
	return &loginHandler{
		uc: usecases.LoginUsecase{
			UserExists:     exists,
			SelectPassword: sel,
			Secret:         sec,
		},
	}
}

// @Summary Login User
// @Description Check if user exists and gives him jwt
// @Tags Users
// @Produce json
// @Success 200 {object} transportation.JWT
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /login [post]
func (login *loginHandler) LoginUser(w http.ResponseWriter, r *http.Request) {

	var cred transportation.Credentials
	if err := json.NewDecoder(r.Body).Decode(&cred); err != nil {
		http.Error(w, "Error decoding credentials", 400)
		return
	}

	jwt, err := login.uc.Exec(transportation.ToDomain(cred))
	if err != nil {

		if errors.Is(err, domain.ErrUserNotExists) {
			http.Error(w, domain.ErrUserNotExists.Error(), 400)
			return
		} else if errors.Is(err, domain.ErrWrongPassword) {
			http.Error(w, domain.ErrWrongPassword.Error(), 400)
			return
		} else if errors.Is(err, domain.ErrCreatingJWT) {
			http.Error(w, domain.ErrCreatingJWT.Error(), 500)
			return
		} else {
			http.Error(w, domain.InternalServerError.Error(), 500)
			return
		}
	}

	if err := json.NewEncoder(w).Encode(transportation.NewJWT(jwt)); err != nil {
		http.Error(w, "Error encoding jwt-token", 500)
		return
	}

}
