package models

import (
	"github.com/jinzhu/gorm"
)

// Post is post models property
type PostReaction struct {
	gorm.Model
	UserID    uint
	PostID    uint
	EmojiCode string
	User      User
	Post      Post
}
