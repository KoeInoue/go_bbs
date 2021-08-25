package models

import (
	"github.com/jinzhu/gorm"
)

// Post is post models property
type PostReaction struct {
	gorm.Model
	ID        uint
	UserId    uint
	PostId    uint
	EmojiCode string
}
