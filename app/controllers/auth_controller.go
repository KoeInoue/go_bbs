package controllers

import (
	"errors"
	"go_bbs/requests"
	"go_bbs/services"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

type AuthController struct{}

// Function to use Validator and Service
func (ac AuthController) PreRegister(c *gin.Context) {
	req := requests.PreRegisterRequest{}
	// Validation
	if isErr := req.ValidatePreRegister(c); isErr {
		return
	}

	// Validate duplicate Email
	if isErr := req.ValidateDuplicateEmail(c); isErr {
		return
	}

	service := services.AuthService{}
	if err := service.PreRegister(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

func (ac AuthController) Register(c *gin.Context) {
	req := requests.RegisterRequest{}
	// Validation
	if isErr := req.ValidateRegister(c); isErr {
		return
	}
	// check email token
	if isErr := req.ValidateEmailToken(c); isErr {
		return
	}

	// register
	service := services.AuthService{}
	if token, err := service.Register(req); err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			respErrors := make(map[string]interface{})
			respErrors["account"] = "account is already registered."
			c.JSON(http.StatusBadRequest, gin.H{"errors": respErrors})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"errors": err})
		return
	} else {
		session := sessions.Default(c)
		session.Set("auth", token)
		session.Save()
		c.JSON(http.StatusCreated, gin.H{
			"status": "ok",
			"token":  token,
		})
		return
	}
}

func (AuthController) Logout(c *gin.Context) {
	//セッションからデータを破棄する
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}
