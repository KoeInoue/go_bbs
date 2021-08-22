package models

import "time"

// Post is post models property
type Post struct {
	ID        uint
	Content   string
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User
}
