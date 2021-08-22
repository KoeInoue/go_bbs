package middleware

import (
	"go_bbs/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header["Token"][0]
		id := c.Request.Header["User-Id"][0]
		repo := repository.AuthRepository{}
		err := repo.VarifyToken(token, id)

		if err != nil {
			c.Status(http.StatusUnauthorized)
			c.Abort()
		} else {
			c.Next()
		}
	}
}
