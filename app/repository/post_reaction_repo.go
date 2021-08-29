package repository

import (
	"go_bbs/db"
	"go_bbs/models"
)

type PostReactionRepository struct{}

func (PostReactionRepository) CreatePostReaction(pr *models.PostReaction) error {
	orm := db.GetDB()
	if err := orm.Create(pr).Error; err != nil {
		return err
	}
	if err := orm.Preload("User").Preload("Post").Find(pr).Error; err != nil {
		return err
	}

	return nil
}
