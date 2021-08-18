package services

import (
	"errors"
	"fmt"
	"go_bbs/mail"
	"go_bbs/models"
	"go_bbs/repository"
	"go_bbs/requests"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

func (AuthService) Register(req requests.RegisterRequest) (string, error) {
	pu := models.PreUser{
		UrlToken: req.Token,
	}

	repo := repository.AuthRepository{}
	pu, err := repo.GetPreUserByToken(pu.UrlToken)
	if err != nil {
		return "", err
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	u := models.User{
		Email:    pu.Email,
		Name:     req.Name,
		Status:   1,
		Password: string(hashed),
	}

	user, err := repo.StoreNewUser(u)
	if err != nil {
		return "", err
	}
	token, err := GenerateTokenProc(fmt.Sprint(user.Id), time.Now())
	if err != nil {
		return "", err
	}

	return token, nil
}

func GenerateTokenProc(userId string, now time.Time) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (AuthService) Login(req requests.LoginRequest) (models.User, string, error) {
	repo := repository.AuthRepository{}
	u, err := repo.GetUserByEmail(req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return u, "", nil
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password))

	if err != nil {
		return u, "", nil
	} else {
		token, err := GenerateTokenProc(fmt.Sprint(u.Id), time.Now())
		if err != nil {
			return u, "err", err
		}

		return u, token, nil
	}
}
