package models

import (
	"github.com/jinzhu/gorm"
)

// Session interface to manage user sessions
type Session struct {
	gorm.Model
	Token  string `json:"token"`
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
}
