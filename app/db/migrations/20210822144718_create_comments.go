package main

import (
	"database/sql"
	"go_bbs/db"
	"time"
)

type Comment struct {
	Id        uint `gorm:"primary_key"`
	UserId    uint
	PostId    uint
	Content   string `gorm:"type:text; not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func init() {
	db.Init()
}

// Up is executed when this migration is applied
func Up_20210822144718(txn *sql.Tx) {
	orm := db.GetDB()
	// Create table for `PreUser`
	orm.CreateTable(&Comment{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE").AddForeignKey("post_id", "posts(id)", "CASCADE", "CASCADE")
}

// Down is executed when this migration is rolled back
func Down_20210822144718(txn *sql.Tx) {
	orm := db.GetDB()
	orm.DropTable(&Comment{})
}
