package domain

import (
	"time"
)

type Users []User

type User struct {
	Id        int
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}
