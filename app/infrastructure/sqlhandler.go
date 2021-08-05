package infrastructure

import (
	"github.com/jinzhu/gorm"

	"viet_college_api/interfaces/database"
)

type SqlHandler struct {
	db  *gorm.DB
	err error
}

func NewSqlHandler() database.SqlHandler {
	Init()
	db := Db()

	sqlHandler := new(SqlHandler)
	sqlHandler.db = db
	return sqlHandler
}

func (handler *SqlHandler) Create(obj interface{}) {
	handler.db.Create(obj)
}

func (handler *SqlHandler) FindAll(obj interface{}) {
	handler.db.Find(obj)
}

func (handler *SqlHandler) DeleteById(obj interface{}, id string) {
	handler.db.Delete(obj, id)
}
