package repository

import (
	"go_bbs/db"
	"go_bbs/models"

	"github.com/jinzhu/gorm"
)

type PostRepository struct{}

type PostResult struct {
	Name    string
	Content string
	UserId  uint
}

func (PostRepository) CreatePost(p *models.Post) error {
	orm := db.GetDB()
	if err := orm.Create(p).Error; err != nil {
		return err
	}
	return nil
}

func (PostRepository) GetPosts(post *[]models.Post) error {
	orm := db.GetDB()
	err := orm.Preload("User").Preload("Comments", func(orm *gorm.DB) *gorm.DB {
		return orm.Order("created_at ASC")
	}).Preload("Comments.User").Preload("PostReactions").Order("created_at DESC").Find(&post).Error
	if err != nil {
		return err
	}

	return nil
}
