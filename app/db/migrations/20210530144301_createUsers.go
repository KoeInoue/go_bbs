package main

import (
	"database/sql"
	"go_bbs/db"
	"time"

	"github.com/jinzhu/gorm"
)

var (
	orm *gorm.DB
	err error
)

type User struct {
	Id        uint   `gorm:"primary_key"`
	Name      string `gorm:"size:255"`
	Email     string `gorm:"size:255"`
	EmailCode string `gorm:"size:255"`
	Password  string `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func init() {
	db.Init()
	orm = db.GetDB()
}

// Up is executed when this migration is applied
func Up_20210530144301(txn *sql.Tx) {
	// Create table for `User`
	orm.CreateTable(&User{})
}

// Down is executed when this migration is rolled back
func Down_20210530144301(txn *sql.Tx) {
	orm.DropTable(&User{})
}
