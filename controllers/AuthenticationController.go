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

type register struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// LoginHandler Login Authentication for users
func LoginHandler(response http.ResponseWriter, request *http.Request) {
	var body login
	json.NewDecoder(request.Body).Decode(&body)

	if body.Email == "" || body.Password == "" {
		services.ErrorResponse(response, http.StatusBadRequest, "Email and Password are required")
	}

	resp, err := services.AuthenticateUser(body.Email, body.Password)
	if err != nil {
		log.Printf("Something went wrong %s", err)
		services.ErrorResponse(response, http.StatusInternalServerError, err.Error())
		return
	}

	services.SuccessResponse(response, "Authentication Successful", resp)
	return
}

// RegisterHandler Registration authentication for users
func RegisterHandler(response http.ResponseWriter, request *http.Request) {
	var body register
	json.NewDecoder(request.Body).Decode(&body)

	// toDo: perform authentication checks
	resp, err := services.RegisterUser(body.FirstName, body.LastName, body.Email, body.Password)
	if err != nil {
		log.Printf("Something went wrong %s", err)
		services.ErrorResponse(response, http.StatusInternalServerError, err.Error())
		return
	}

	services.SuccessResponse(response, "Registration Successful", resp)
	return
}
