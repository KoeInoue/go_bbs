package repository

import (
	"fmt"
	"go_bbs/db"
	"go_bbs/models"
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
	err := orm.Preload("User").Preload("Comments").Preload("PostReactions").Order("created_at desc").Find(&post)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}
