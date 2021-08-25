package models

import (
	"github.com/jinzhu/gorm"
)

// Post is post models property
type Comment struct {
	gorm.Model
	ID      uint
	UserId  uint
	PostId  uint
	Content string
}
