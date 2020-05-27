package models

import "github.com/jinzhu/gorm"

var db *gorm.DB

// SetDB make the connection available in the model scope
func SetDB(conn *gorm.DB) {
	db = conn
}
