package services

import (
	"go_bbs/models"
	"go_bbs/repository"
	"go_bbs/requests"
)

type CommentService struct{}

func (CommentService) CreateComment(req requests.CommentRequest) (models.Comment, error) {
	comment := models.Comment{
		Content: req.Content,
		UserID:  req.UserID,
		PostID:  req.PostID,
	}

	repo := repository.CommentRepository{}
	if err := repo.CreateComment(&comment); err != nil {
		return comment, err
	}

	return comment, nil
}
