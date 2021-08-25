package models

import (
	"github.com/jinzhu/gorm"
)

// Post is post models property
type Post struct {
	gorm.Model
	ID            uint
	Content       string
	UserID        uint
	User          User
	Comments      []Comment
	PostReactions []PostReaction
}
