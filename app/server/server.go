package server

import (
	"go_bbs/controllers"

	"github.com/gin-gonic/gin"
)

// Init is initialize server
func Init() *gin.Engine {
	return router()
}

func router() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	{
		authCtrl := controllers.AuthController{}
		api.POST("register", authCtrl.Register)
	}

	return r
}
