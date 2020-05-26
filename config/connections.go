package config

import (
	"time"

	"github.com/jinzhu/gorm"

	// postgres db plugin
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB representing database connection instance
var db *gorm.DB

// OpenDatabaseConnection a DB connection
func OpenDatabaseConnection(username, password, dbName, dbHost, dbPort, sslmode string) (*gorm.DB, error) {
	var err error
	db, err = gorm.Open("mysql", username+":"+password+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		return nil, err
	}

	db.DB().SetConnMaxLifetime(time.Minute * 5)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(10)
	return db, nil
}
