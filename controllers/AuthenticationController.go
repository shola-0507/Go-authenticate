package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Go-authenticate/services"
)

type login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type register struct{}

// LoginHandler Login Authentication for users
func LoginHandler(response http.ResponseWriter, request *http.Request) {
	var body login
	json.NewDecoder(request.Body).Decode(&body)

	if body.Email == "" || body.Password == "" {
		services.ErrorResponse(response, http.StatusBadRequest, "Email and Password are required")
		return
	}

	resp, err := services.AuthenticateUser(body.Email, body.Password)
	if err != nil {
		log.Printf("Something went wrong %s", err)
		services.ErrorResponse(response, http.StatusInternalServerError, "Something went wrong. Please try again later")
		return
	}

	services.SuccessResponse(response, "Authentication Successful", resp)
	return
}

// RegisterHandler Registration authentication for users
func RegisterHandler(response http.ResponseWriter, request *http.Request) {}
