package services

import (
	"fmt"
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
	if err := repo.GetPosts(&posts); err != nil {
		return posts, nil
	}

	return posts, nil
}

func (PostService) DeletePost(req requests.PostIdRequest) error {
	repo := repository.PostRepository{}
	id := req.ID
	if err := repo.DeletePost(&id); err != nil {
		return err
	}

	return nil
}

func (PostService) EditPost(req requests.PostEditRequest) error {
	repo := repository.PostRepository{}
	id := req.ID
	content := req.Content
	fmt.Println(req)
	fmt.Println("a")

	if err := repo.EditPost(id, content); err != nil {
		return err
	}

	return nil
}
