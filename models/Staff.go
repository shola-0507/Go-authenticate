package models

import (
	"github.com/jinzhu/gorm"
)

// Staff struct defines the user model interface
type Staff struct {
	gorm.Model
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Email         string `json:"email"`
	AccountNumber string `json:"account_number"`
	JobTitle      string `json:"job_title"`
	Password      string `json:"password"`
	CompanyID     int    `json:"company_id"`
}
