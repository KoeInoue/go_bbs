package repository

import (
	"errors"
	"go_bbs/db"
	"go_bbs/models"

	"github.com/jinzhu/gorm"
)

type AuthRepository struct{}

func (AuthRepository) CreatePreUser(u *models.PreUser) error {
	orm := db.GetDB()
	if err := orm.Create(u).Error; err != nil {
		return err
	}
	return nil
}

func (AuthRepository) GetPreUserByEmail(email string) (string, error) {
	orm := db.GetDB()
	pu := models.PreUser{}
	if err := orm.Table("pre_users").Where("email = ?", email).First(&pu).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", nil
		}
		return "", err
	}

	return pu.Email, nil
}

func (AuthRepository) IsPreUserExists(token string) (bool, error) {
	orm := db.GetDB()
	pu := models.PreUser{}
	if err := orm.Table("pre_users").Where("url_token = ?", token).First(&pu).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (AuthRepository) GetPreUserByToken(token string) (models.PreUser, error) {
	orm := db.GetDB()
	pu := models.PreUser{}
	if err := orm.Table("pre_users").Where("url_token = ?", token).First(&pu).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return pu, nil
		}
		return pu, err
	}

	return pu, nil
}

func (AuthRepository) StoreNewUser(user models.User) (models.User, error) {
	orm := db.GetDB()
	if result := orm.Create(&user); result.Error != nil {
		return user, result.Error
	} else {
		return user, nil
	}
}
