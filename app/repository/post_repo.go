package repository

import (
	"go_bbs/db"
	"go_bbs/models"
)

type PostRepository struct{}

func (PostRepository) CreatePost(p *models.Post) error {
	orm := db.GetDB()
	if err := orm.Create(p).Error; err != nil {
		return err
	}
	return nil
}

func (PostRepository) GetPost(posts *[]models.Post) error {
	orm := db.GetDB()
	if err := orm.Table("posts").Order("created_at").Find(&posts).Error; err != nil {
		return err
	}

	return nil
}

func (PostRepository) GetPostRelation(post *models.Post, userId uint) error {
	orm := db.GetDB()
	if err := orm.Table("users").Select("name").Where("id = ?", userId).Find(&post.User).Error; err != nil {
		return err
	}

	return nil
}
