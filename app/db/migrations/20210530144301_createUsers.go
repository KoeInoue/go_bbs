package main

import (
	"database/sql"
	"go_bbs/db"
	"time"
)

type User struct {
	Id        uint         `gorm:"primary_key"`
	Name      string       `gorm:"size:255; not null"`
	Email     string       `gorm:"unique; size:255; not null"`
	Status    sql.NullBool `gorm:"default:false"`
	Password  string       `gorm:"size:255; not null"`
	Token     string       `gorm:"unique; size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func init() {
	db.Init()
}

// Up is executed when this migration is applied
func Up_20210530144301(txn *sql.Tx) {
	orm := db.GetDB()
	// Create table for `User`
	orm.CreateTable(&User{})
}

// Down is executed when this migration is rolled back
func Down_20210530144301(txn *sql.Tx) {
	orm := db.GetDB()
	orm.DropTable(&User{})
}
