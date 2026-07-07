package handlers

import (
	"auth/internal/domain"
	"auth/internal/interfaces"
	"auth/internal/transportation"
	"auth/internal/usecases"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type registerHandler struct {
	uc usecases.RegisterUsecase
}

func NewRegisterHandler(
	exists interfaces.UserExistsInterface,
	insert interfaces.InsertCredentialsInterface,
) *registerHandler {
	return &registerHandler{
		uc: usecases.RegisterUsecase{
			UserExists:        exists,
			InsertCredentials: insert,
		},
	}
}

// @Summary Registrate user
// @Description Inserts new login and password into table users, if login doesn't already exist
// @Tags Users
// @Produce json
// @Success 200 {object} transportation.Credentials
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Failure 502 {string} string
// @Router /register [post]
func (reg *registerHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {

	var credHTTP transportation.Credentials
	if err := json.NewDecoder(r.Body).Decode(&credHTTP); err != nil {
		http.Error(w, "Error decoding credentials", 400)
		return
	}

	credDomain := transportation.ToDomain(credHTTP)

	if err := domain.Validating(credDomain); err != nil {

		log.Println("Validating error:", err)
		if errors.Is(err, domain.ErrForbiddenSymbols) {
			http.Error(w, domain.ErrForbiddenSymbols.Error(), 400)
		} else if errors.Is(err, domain.ErrLogin) {
			http.Error(w, domain.ErrLogin.Error(), 400)
		} else if errors.Is(err, domain.ErrPassword) {
			http.Error(w, domain.ErrPassword.Error(), 400)
		}
		return
	}

	domainCred, err := reg.uc.Exec(credDomain)
	if err != nil {

		if errors.Is(err, domain.ErrUserExists) {
			http.Error(w, domain.ErrUserExists.Error(), 400)
			return
		}

		log.Println("Error in usecase during registrating:", err)
		http.Error(w, "Internal server error", 502)
		return
	}

	if err = json.NewEncoder(w).Encode(transportation.ToHttp(*domainCred)); err != nil {
		http.Error(w, "Error encoding credentials", 500)
		return
	}

}
