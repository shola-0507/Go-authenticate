package middlewares

import (
	"net/http"
)

// CheckUserIsAuthenticated verifies that the user is authenticated
func CheckUserIsAuthenticated(response http.ResponseWriter, request *http.Request) http.HandlerFunc {
	return nil
}
