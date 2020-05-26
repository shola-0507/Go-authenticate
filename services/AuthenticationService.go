package services

import (
	"errors"
	"log"

	"github.com/Go-authenticate/models"
)

// AuthenticateUser Handles user authentication
func AuthenticateUser(email, password string) (interface{}, error) {
	response := make(map[string]interface{})

	user, err := new(models.User).FindUserByEmail(email)

	if err != nil {
		log.Printf("Something went wrong %s", err)
		return nil, errors.New("Email and Password mismatch")
	}

	hashedPwd, err := EncryptPassword([]byte(password))
	if err != nil {
		return nil, err
	}

	passwordMatches := ComparePasswords(hashedPwd, []byte(user.Password))
	if !passwordMatches {
		return nil, errors.New("Email and Password mismatch")
	}

	token, err := GenerateJWT(email, user.Role)
	if err != nil {
		return response, err
	}

	response["token_type"] = "Bearer"
	response["jwt_token"] = token["token"]
	response["expires_at"] = token["expiry"]
	return response, nil
}

// RegisterUser Handles user registeration
func RegisterUser(firstName, lastName, email, password string) (interface{}, error) {
	response := make(map[string]interface{})
	hashedPwd, err := EncryptPassword([]byte(password))
	if err != nil {
		return nil, err
	}

	newUser := &models.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  hashedPwd,
		Role:      "user",
	}

	if err := newUser.Create(); err != nil {
		return nil, errors.New("Something went wrong registering the user")
	}

	response["first_name"] = newUser.FirstName
	response["last_name"] = newUser.LastName
	response["email"] = newUser.Email
	return response, nil
}
