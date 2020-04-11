package services

// AuthenticateUser Handles user authentication
func AuthenticateUser(email, password string) (map[string]interface{}, error) {
	// make database call and get user credentials
	role := "Admin"
	response := make(map[string]interface{})
	token, err := GenerateJWT(email, role)
	if err != nil {
		return response, err
	}
	response["token_type"] = "Bearer"
	response["jwt_token"] = token["token"]
	response["expires_at"] = token["expiry"]

	return response, nil
}
