package config

import (
	"github.com/jinzhu/gorm"

	// postgres db plugin
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB representing database connection instance
var DB *gorm.DB

// OpenDatabaseConnection a DB connection
func OpenDatabaseConnection(user, password, dbname, host, port, sslmode string) error {
	var err error
	DB, err = gorm.Open("postgres", "host="+host+" port="+port+" user="+user+" dbname="+dbname+" password="+password+" sslmode="+sslmode)

	if err != nil {
		return err
	}

	return nil
}

// CloseDatabaseConnection a DB connection
func CloseDatabaseConnection() error {
	return DB.Close()
}
