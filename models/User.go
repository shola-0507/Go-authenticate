package models

import (
	"github.com/Go-authenticate/config"
	"github.com/jinzhu/gorm"
)

// User struct defines the user model interface
type User struct {
	gorm.Model
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Email         string `json:"email"`
	AccountNumber string `json:"account_number"`
	JobTitle      string `json:"job_title"`
	Password      string `json:"password"`
	CompanyID     int    `json:"company_id"`
	Salary        int    `json:"salary"`
	Role          string `json:"role"`
}

var db = config.GetDatabase()

// FindUser record in the DB
func FindUser(email string) (User, error) {
	var user User

	if err := db.Where(&User{Email: email}).Find(&user).Error; err != nil {
		return User{}, err
	}

	return user, nil
}
