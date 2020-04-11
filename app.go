package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Go-authenticate/config"
	"github.com/gorilla/mux"
)

// App application App interface
type App struct {
	Router *mux.Router
}

// Initialize the go Application with Database setup
func (app *App) Initialize(user, password, dbname, host, port, sslmode string) {
	if err := config.OpenDatabaseConnection(user, password, dbname, host, port, sslmode); err != nil {
		log.Fatalf("Something went wrong connecting to the database %s", err)
	}

	defer config.CloseDatabaseConnection()
	app.Router = mux.NewRouter()
	app.initializeRoutes()
}

// Run the go app
func (app *App) Run(address string) {
	server := &http.Server{
		Handler:      app.Router,
		Addr:         address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Starting application on.... " + address)
	log.Fatal(server.ListenAndServe())
}
