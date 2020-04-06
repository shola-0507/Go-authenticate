package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// App application App interface
type App struct {
	Router *mux.Router
}

// Initialize the go Application with Database setup
func (app *App) Initialize(user, password, dbname, host, port string) {
	db, err := gorm.Open("postgres", "host="+host+"port="+port+"user="+user+"dbname="+dbname+"password="+password)
	defer db.Close()

	if err != nil {
		log.Fatalf("Something went wrong connecting to the database %s", err)
	}

	app.Router = mux.NewRouter()
	app.initializeRoutes()
}

// Run the go app
func (app *App) Run(address string) {
	log.Fatal(http.ListenAndServe(address, app.Router))
}
