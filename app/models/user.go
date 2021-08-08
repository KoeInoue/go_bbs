package models

import "time"

// swagger:model User
type User struct {
	ID        uint
	Name      string
	Email     string
	MailCode  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
