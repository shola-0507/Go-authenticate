package models

import (
	"github.com/jinzhu/gorm"
)

// User struct defines the user model interface
type User struct {
	gorm.Model
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
	RoleID      uint   `json:"role_id"`
	Role        Role
}

// TableName set table name for model
func (User) TableName() string {
	return "users"
}

// FindUserByEmail find user record by email
func (user *User) FindUserByEmail(email string) (*User, error) {
	var result User
	if err := db.Preload("Role").Find(&result, User{Email: email}).Error; err != nil {
		return &result, err
	}

	return &result, nil
}

// Create new user record
func (user *User) Create() error {
	if err := db.Create(user).Error; err != nil {
		return err
	}

	return nil
}
