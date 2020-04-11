package services

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
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
