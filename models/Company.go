package models

import (
	"github.com/jinzhu/gorm"
)

// Company model interface
type Company struct {
	gorm.Model
	Name string `json:"name"`
}
