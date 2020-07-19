package services

import (
	"os"
	"strconv"
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
func GenerateJWT(email, role string) (map[string]string, error) {
	response := make(map[string]string)
	jwtKey := []byte(os.Getenv("APP_JWT_KEY"))
	jwtExpiryTime, err := strconv.Atoi(os.Getenv("APP_JWT_EXPIRY"))
	if err != nil {
		return response, err
	}

	expirationTime := time.Now().Add(time.Duration(jwtExpiryTime) * time.Second)
	claims := &claims{
		Email: email,
		Role:  role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return response, err
	}
	response["token"] = tokenString
	response["expiry"] = expirationTime.UTC().String()
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
func ComparePasswords(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	bytePassword := []byte(plainPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, bytePassword)
	if err != nil {
		return false
	}

	return true
}
