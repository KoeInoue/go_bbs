package models

import "time"

type User struct {
	Id        int
	Name      string
	Email     string
	Status    int
	Password  string
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
