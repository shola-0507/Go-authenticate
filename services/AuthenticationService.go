package services

import (
	"errors"

	"github.com/Go-authenticate/models"
)

// AuthenticateUser Handles user authentication
func AuthenticateUser(email, password string) (map[string]interface{}, error) {
	response := make(map[string]interface{})
	user, err := models.FindUser(email)
	if err != nil {
		return nil, errors.New("Username or password doesn't exist")
	}

	hashedPwd, err := EncryptPassword([]byte(password))

	if err != nil {
		return nil, err
	}

	passwordMatches := ComparePasswords(hashedPwd, []byte(password))

	if !passwordMatches {
		return nil, errors.New("Password Mismatch")
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
