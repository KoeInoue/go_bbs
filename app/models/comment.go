package models

import (
	"github.com/jinzhu/gorm"
)

// Post is post models property
type Comment struct {
	gorm.Model
	UserID  uint
	PostID  uint
	Content string
	User    User
}
