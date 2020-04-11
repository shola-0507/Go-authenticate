package main

import (
	"encoding/json"
	"net/http"

	"github.com/Go-authenticate/controllers"
)

func (app *App) initializeRoutes() {
	apiV1 := app.Router.PathPrefix("/api/v1").Subrouter()
	app.Router.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		resp := map[string]string{
			"version": "1.0.0",
			"status":  "running",
		}
		response.Header().Set("Content-Type", "Application/json")
		json.NewEncoder(response).Encode(resp)
	}).Methods("GET")
	apiV1.HandleFunc("/login", controllers.LoginHandler).Methods("POST")
	apiV1.HandleFunc("/register", controllers.RegisterHandler).Methods("POST")
}
