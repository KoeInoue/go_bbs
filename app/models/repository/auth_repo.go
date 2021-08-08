package repository

import (
	"go_bbs/db"
	"go_bbs/models"

	"github.com/gin-gonic/gin"
)

// Service procides user's behavior
type AuthRepository struct{}

// User is alias of entity.User struct
type User models.User

// CreateModel is create User model
func (_ AuthRepository) CreateUser(c *gin.Context) (User, error) {
	db := db.GetDB()
	u := User{}
	if err := c.Bind(&u); err != nil {
		return u, err
	}

	if err := db.Create(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}
