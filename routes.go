package main

import (
	"encoding/json"
	"net/http"

	"github.com/Go-authenticate/controllers"
)

func (app *App) initializeRoutes() {
	app.Router.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Content-Type", "Application/json")
		json.NewEncoder(response).Encode(map[string]string{"Application Status": "Running"})
	}).Methods("GET")
	app.Router.HandleFunc("api/v1/login", controllers.LoginHandler).Methods("GET")
}
