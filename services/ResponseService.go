package services

import (
	"encoding/json"
	"net/http"
)

type err struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type success struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ErrorResponse method retuns an http error response
func ErrorResponse(response http.ResponseWriter, status int, message string) {
	data := &err{
		Status:  "failed",
		Message: message,
	}

	response.Header().Set("Content-Type", "Application/json")
	response.WriteHeader(status)
	json.NewEncoder(response).Encode(data)
	return
}

// SuccessResponse method retuns an http success response
func SuccessResponse(response http.ResponseWriter, message string, payload interface{}) {
	data := &success{
		Status:  "success",
		Message: message,
		Data:    payload,
	}

	response.Header().Set("Content-Type", "Application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(data)
	return
}
