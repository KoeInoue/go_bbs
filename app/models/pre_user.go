package models

import (
	"database/sql"

	"github.com/jinzhu/gorm"
)

// swagger:model User
type PreUser struct {
	gorm.Model
	ID       uint
	UrlToken string
	Email    string
	Flag     sql.NullBool
}
