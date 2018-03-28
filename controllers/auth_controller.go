package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/quakkels/goalie/models"
	"github.com/quakkels/goalie/services"
)

// Login logs in a user
func Login(w http.ResponseWriter, r *http.Request) {
	requestUser := new(models.User)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestUser)

	responseStatus, token := services.Login(requestUser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
	w.Write(token)
}

func RefreshToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {}

func Logout(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {}
