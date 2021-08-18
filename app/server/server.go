package server

import (
	"go_bbs/controllers"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) *gin.Engine {
	api := r.Group("/api")
	{
		authCtrl := controllers.AuthController{}
		api.POST("pre-register", authCtrl.PreRegister)
		api.POST("register", authCtrl.Register)
		api.POST("login", authCtrl.Login)
		api.GET("logout", authCtrl.Logout)
	}

	return r
}
