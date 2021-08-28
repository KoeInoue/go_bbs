package repository

import (
	"go_bbs/db"
	"go_bbs/models"
)

type CommentRepository struct{}

func (CommentRepository) CreateComment(c *models.Comment) error {
	orm := db.GetDB()
	if err := orm.Create(c).Preload("User").Error; err != nil {
		return err
	}
	if err := orm.Preload("User").Find(c).Error; err != nil {
		return err
	}

	return nil
}
