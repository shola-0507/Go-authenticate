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

	passwordMatches := ComparePasswords(user.Password, password)
	if !passwordMatches {
		return nil, errors.New("Email and Password mismatch")
	}

	jwtInfo, err := GenerateJWT(email, user.Role.Name)
	if err != nil {
		return response, err
	}

	// create user session
	newSession := &models.Session{
		Token:     jwtInfo["token"],
		UserID:    int(user.ID),
		ExpiredAt: jwtInfo["expiry"],
	}

	if err := newSession.Create(); err != nil {
		return nil, errors.New("Something went wrong registering the user")
	}

	response["token_type"] = "Bearer"
	response["jwt_token"] = jwtInfo["token"]
	response["expires_at"] = jwtInfo["expiry"]
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
