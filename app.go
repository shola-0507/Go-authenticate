package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Go-authenticate/config"
	"github.com/Go-authenticate/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// App application App interface
type App struct {
	Router   *mux.Router
	Database *gorm.DB
}

// Initialize the go Application with Database setup
func (app *App) Initialize(user, password, dbname, host, port, sslmode string) {
	db, err := config.OpenDatabaseConnection(user, password, dbname, host, port, sslmode)

	if err != nil {
		log.Fatalf("Something went wrong connecting to the database %s", err)
	}

	app.Database = db
	models.SetDB(app.Database)
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
	log.Println("Application running on.... " + address)
	log.Fatal(server.ListenAndServe())
}

// Close the database connection
func (app *App) Close() error {
	return app.Database.Close()
}
