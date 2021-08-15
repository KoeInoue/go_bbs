package models

import "time"

// swagger:model User
type User struct {
	Id        int
	Name      string
	Email     string
	Status    int
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
