package main

import (
	"database/sql"
	"go_bbs/db"
	"time"
)

type PreUser struct {
	Id        uint         `gorm:"primary_key"`
	UrlToken  string       `gorm:"unique; size:255; not null"`
	Email     string       `gorm:"unique; size:255; not null"`
	Flag      sql.NullBool `gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func init() {
	db.Init()
}

// Up is executed when this migration is applied
func Up_20210811132223(txn *sql.Tx) {
	orm := db.GetDB()
	// Create table for `PreUser`
	orm.CreateTable(&PreUser{})
}

// Down is executed when this migration is rolled back
func Down_20210811132223(txn *sql.Tx) {
	orm := db.GetDB()
	orm.DropTable(&PreUser{})
}
