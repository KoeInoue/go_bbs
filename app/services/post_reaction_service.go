package services

import (
	"go_bbs/models"
	"go_bbs/repository"
	"go_bbs/requests"
)

type PostReactionService struct{}

func (PostReactionService) CreatePostReaction(req requests.PostReactionRequest) (models.PostReaction, error) {
	pr := models.PostReaction{
		EmojiCode: req.EmojiCode,
		UserID:    req.UserID,
		PostID:    req.PostID,
	}

	repo := repository.PostReactionRepository{}
	if err := repo.CreatePostReaction(&pr); err != nil {
		return pr, err
	}

	return pr, nil
}

func (PostReactionService) DeletePostReaction(req requests.PostReactionIdRequest) error {
	repo := repository.PostReactionRepository{}
	id := req.ID
	if err := repo.DeletePostReaction(&id); err != nil {
		return err
	}

	return nil
}
