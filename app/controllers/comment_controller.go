package controllers

import (
	"go_bbs/requests"
	"go_bbs/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommentController struct{}

func (CommentController) CreateComment(c *gin.Context) {
	req := requests.CommentRequest{}
	// Validation
	if isErr := req.ValidateComment(c); isErr {
		return
	}

	service := services.CommentService{}
	if comment, err := service.CreateComment(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": err})
		return
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"status":  "ok",
			"comment": comment,
		})
	}

}
