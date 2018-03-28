package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/quakkels/goalie/authentication"
	"github.com/quakkels/goalie/models"
)

// Login the user
func Login(requestUser *models.User) (int, []byte) {
	authBackend := authentication.InitJWTAuthenticationBackend()
	if authBackend.Authenticate(requestUser) {
		token, err := authBackend.GenerateToken(requestUser.UUID)
		if err != nil {
			fmt.Println("Error making token.")
			return http.StatusInternalServerError, nil
		}
		response, _ := json.Marshal(models.TokenAuthentication{Token: token})
		return http.StatusOK, response

	}

	return http.StatusUnauthorized, nil
}
