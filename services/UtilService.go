package services

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type claims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

// GenerateJWT generate a new JWT for the user
func GenerateJWT(email, role string) (map[string]interface{}, error) {
	response := make(map[string]interface{})
	jwtKey := []byte(os.Getenv("APP_JWT_KEY"))
	expirationTime := time.Now().Add(24 * time.Hour).Unix()
	claims := &claims{
		Email: email,
		Role:  role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return response, err
	}
	response["token"] = tokenString
	response["expiry"] = expirationTime
	return response, nil
}

// EncryptPassword encrypt users password
func EncryptPassword(password []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// ComparePasswords Check to see that the users password matches the encrypted record
func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false
	}

	return true
}
