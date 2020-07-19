package models

import "github.com/jinzhu/gorm"

// Session interface to manage user sessions
type Session struct {
	gorm.Model
	Token     string `json:"token"`
	UserID    int    `json:"user_id"`
	ExpiredAt string `json:"expired_at"`
}

// TableName set table name for model
func (Session) TableName() string {
	return "sessions"
}

// Create add user record to the DB
func (session *Session) Create() error {
	if err := db.Create(session).Error; err != nil {
		return err
	}

	return nil
}
