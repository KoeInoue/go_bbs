package models

import (
	"database/sql"
	"time"
)

// swagger:model User
type PreUser struct {
	Id        uint
	UrlToken  string
	Email     string
	Flag      sql.NullBool
	CreatedAt time.Time
	UpdatedAt time.Time
}
