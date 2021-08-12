package repository

import (
	"errors"
	"go_bbs/db"
	"go_bbs/models"

	"github.com/jinzhu/gorm"
)

type AuthRepository struct{}

func (AuthRepository) CreatePreUser(u *models.PreUser) error {
	db := db.GetDB()

	if err := db.Create(u).Error; err != nil {
		return err
	}
	return nil
}

func (AuthRepository) GetPreUserByEmail(email string) (string, error) {
	db := db.GetDB()
	pu := models.PreUser{}
	if err := db.Table("pre_users").Where("email = ?", email).First(&pu).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", nil
		}
		return "", err
	}

	return pu.Email, nil
}
