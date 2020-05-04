package config

import (
	"github.com/jinzhu/gorm"

	// postgres db plugin
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB representing database connection instance
var DB *gorm.DB

// OpenDatabaseConnection a DB connection
func OpenDatabaseConnection(username, password, dbName, dbHost, dbPort, sslmode string) error {
	var err error
	DB, err = gorm.Open("mysql", username+":"+password+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		return err
	}

	return nil
}

// CloseDatabaseConnection a DB connection
func CloseDatabaseConnection() error {
	return DB.Close()
}

// GetDatabase returns an instance of the database connection
func GetDatabase() *gorm.DB {
	return DB
}
