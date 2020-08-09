package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Session interface to manage user sessions
type Session struct {
	gorm.Model
	Token     string    `json:"token"`
	UserID    int       `json:"user_id"`
	ExpiresAt time.Time `json:"expires_at"`
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

// FindActiveSession find the most recently created session
func (session *Session) FindActiveSession(userID int) (*Session, error) {
	var result Session
	currentTime := time.Now().Format(time.RFC3339)

	if err := db.Where(&Session{UserID: userID}).Where("expires_at > ?", currentTime).Limit(1).Find(&result).Error; err != nil {
		return &result, err
	}

	return &result, nil
}
