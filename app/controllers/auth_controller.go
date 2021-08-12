package controllers

import (
	"go_bbs/requests"
	"go_bbs/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct{}

// Function to use Validator and Service
func (ac AuthController) PreRegister(c *gin.Context) {
	req := requests.PreRegisterRequest{}
	// Validation
	if isErr := req.Validate(c); isErr {
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
