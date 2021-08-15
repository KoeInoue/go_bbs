package requests

import "time"

// swagger:model User
type User struct {
	Id        uint   `gorm:"primary_key" json:"id" binding:"required"`
	Name      string `gorm:"size:255" json:"name" binding:"required"`
	Email     string `gorm:"size:255"`
	Password  string `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
