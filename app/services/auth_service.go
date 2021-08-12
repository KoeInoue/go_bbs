package services

import (
	"go_bbs/mail"
	"go_bbs/models"
	"go_bbs/repository"
	"go_bbs/requests"

	"github.com/google/uuid"
)

type AuthService struct{}

// Method to create new pre user and send email to the pre user
func (AuthService) PreRegister(req requests.PreRegisterRequest) error {
	u := models.PreUser{
		Email:    req.Email,
		UrlToken: uuid.New().String(),
	}

	repo := repository.AuthRepository{}
	if err := repo.CreatePreUser(&u); err != nil {
		return err
	}

	if err := mail.NewMailer().Send(u.UrlToken); err != nil {
		return err
	}

	return nil
}
