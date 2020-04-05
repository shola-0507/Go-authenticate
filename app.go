package main

import (
	"database/sql"

	"github.com/gorilla/mux"
)

// App application App interface
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize the go Application with Database setup
func (a *App) Initialize(user, password, dbname string) {}

// Run the go app
func (a *App) Run(address string) {}
