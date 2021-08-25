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

func (AuthRepository) GetUserByEmail(email string) (models.User, error) {
	orm := db.GetDB()
	u := models.User{}
	if err := orm.Table("users").Where("email = ?", email).Find(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

func (AuthRepository) StoreUserToken(user models.User, token string) error {
	orm := db.GetDB()
	user.Token = token

	if err := orm.Model(&user).Update("Token", token).Error; err != nil {
		return err
	}

	return nil
}

func (AuthRepository) VarifyToken(token string, id string) error {
	orm := db.GetDB()
	u := models.User{}
	if err := orm.Where("token = ?", token).First(&u, id).Error; err != nil {
		return err
	}

	return nil
}
