package models

import "time"

// Post is post models property
type Post struct {
	ID        uint   `json:"id" binding:"required"`
	Content   string `json:"content" binding:"required"`
	UserID    uint   `gorm:"not null"  json:"user_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
