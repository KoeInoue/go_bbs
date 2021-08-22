package controllers

import (
	"go_bbs/requests"
	"go_bbs/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostController struct{}

func (PostController) CreatePost(c *gin.Context) {
	req := requests.PostRequest{}
	// Validation
	if isErr := req.ValidatePost(c); isErr {
		return
	}

	service := services.PostService{}
	service.CreatePost(req)

	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

func (PostController) GetPosts(c *gin.Context) {
	service := services.PostService{}
	posts, err := service.GetPosts()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
		"posts":  posts,
	})
}
