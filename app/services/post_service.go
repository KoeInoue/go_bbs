package services

import (
	"go_bbs/models"
	"go_bbs/repository"
	"go_bbs/requests"
)

type PostService struct{}

func (PostService) CreatePost(req requests.PostRequest) error {
	p := models.Post{
		Content: req.Content,
		UserID:  req.UserID,
	}

	repo := repository.PostRepository{}
	if err := repo.CreatePost(&p); err != nil {
		return err
	}

	return nil
}

func (PostService) GetPosts() ([]models.Post, error) {
	repo := repository.PostRepository{}

	var posts []models.Post
	if err := repo.GetPost(&posts); err != nil {
		return posts, nil
	}

	// fmt.Println(posts)

	return posts, nil
}
