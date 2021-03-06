package controllers

import (
	"go_bbs/requests"
	"go_bbs/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostReactionController struct{}

func (PostReactionController) CreatePostReaction(c *gin.Context) {
	req := requests.PostReactionRequest{}
	// Validation
	if isErr := req.ValidatePostReaction(c); isErr {
		return
	}

	service := services.PostReactionService{}
	if reaction, err := service.CreatePostReaction(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": err})
		return
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"status":   "ok",
			"reaction": reaction,
		})
	}
}

func (PostReactionController) DeletePostReaction(c *gin.Context) {
	var req requests.PostReactionIdRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": err})
		return
	}

	service := services.PostReactionService{}
	if err := service.DeletePostReaction(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": err})
		return
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"status": "ok",
		})
	}
}
