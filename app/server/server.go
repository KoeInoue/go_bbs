package server

import (
	"go_bbs/controllers"
	"go_bbs/middleware"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) *gin.Engine {
	api := r.Group("/api")
	{
		authCtrl := controllers.AuthController{}
		api.POST("pre-register", authCtrl.PreRegister)
		api.POST("register", authCtrl.Register)
		api.POST("login", authCtrl.Login)
		api.Use(middleware.AuthMiddleware())
		{
			postCtrl := controllers.PostController{}
			api.POST("create-post", postCtrl.CreatePost)
			api.GET("get-posts", postCtrl.GetPosts)
			commCtrl := controllers.CommentController{}
			api.POST("create-comment", commCtrl.CreateComment)
		}
	}

	return r
}
