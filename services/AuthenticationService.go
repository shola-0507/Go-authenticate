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
func RegisterUser(firstName, lastName, email, phoneNumber, password string) (interface{}, error) {
	hashedPassword, err := EncryptPassword([]byte(password))
	if err != nil {
		return nil, err
	}

	userInstance := &models.User{
		FirstName:   firstName,
		LastName:    lastName,
		Email:       email,
		PhoneNumber: phoneNumber,
		Password:    hashedPassword,
		RoleID:      1,
	}

	user, err := new(models.User).FindUserByEmail(email)
	if err != nil {
		log.Printf("Something went wrong %s", err)
		return nil, errors.New("Something went wrong registering the user")
	}

	if user != nil {
		return nil, errors.New("This email address belongs to an existing user")
	}

	if err := userInstance.Create(); err != nil {
		log.Printf("Something went wrong %s", err)
		return nil, errors.New("Something went wrong registering the user")
	}

	response := make(map[string]interface{})
	response["first_name"] = userInstance.FirstName
	response["last_name"] = userInstance.LastName
	response["email"] = userInstance.Email
	return response, nil
}
