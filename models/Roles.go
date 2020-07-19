package models

import "github.com/jinzhu/gorm"

// Role struct to manage user roles and permissions
type Role struct {
	gorm.Model
	Name        string `json:"name"`
	Permissions string `json:"permissions"`
}

// TableName set table name for model
func (Role) TableName() string {
	return "roles"
}
