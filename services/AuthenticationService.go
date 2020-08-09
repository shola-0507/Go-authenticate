package services

import (
	"errors"
	"log"
	"time"

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

	passwordMatches := ComparePasswords(user.Password, password)
	if !passwordMatches {
		return nil, errors.New("Email and Password mismatch")
	}

	// check if user has an unexpired token
	session, err := new(models.Session).FindActiveSession(int(user.ID))
	if *session != (models.Session{}) {
		response["token_type"] = "Bearer"
		response["jwt_token"] = session.Token
		response["expires_at"] = session.ExpiresAt.Format(time.RFC3339)

		return response, nil
	}

	jwtInfo, err := GenerateJWT(email, user.Role.Name)
	if err != nil {
		return response, err
	}

	expiresAt, err := time.Parse(time.RFC3339, jwtInfo["expires_at"])
	if err != nil {
		return response, err
	}

	newSession := &models.Session{
		Token:     jwtInfo["token"],
		UserID:    int(user.ID),
		ExpiresAt: expiresAt,
	}

	if err := newSession.Create(); err != nil {
		return nil, errors.New("Something went wrong registering the user")
	}

	response["token_type"] = "Bearer"
	response["jwt_token"] = jwtInfo["token"]
	response["expires_at"] = jwtInfo["expires_at"]
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
		RoleID:    1,
	}

	if err := newUser.Create(); err != nil {
		return nil, errors.New("Something went wrong registering the user")
	}

	response["first_name"] = newUser.FirstName
	response["last_name"] = newUser.LastName
	response["email"] = newUser.Email
	return response, nil
}
