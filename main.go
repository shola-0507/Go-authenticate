package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file")
	}
	app := App{}
	var address string
	app.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
		os.Getenv("APP_DB_HOST"),
		os.Getenv("APP_DB_PORT"),
		os.Getenv("APP_DB_SSLMODE"))

	hostAddress := os.Getenv("APP_HOST_ADDRESS")
	portNumber := os.Getenv("APP_PORT")
	address = hostAddress + ":" + portNumber
	app.Run(address)
}
