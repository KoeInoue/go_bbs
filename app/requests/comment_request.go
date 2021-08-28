package requests

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Post is post models property
type CommentRequest struct {
	Content string `json:"content" binding:"required"`
	UserID  uint   `json:"userId" binding:"required"`
	PostID  uint   `json:"postId" binding:"required"`
}

// Function to validate that it match the rule of PreRegisterRequest
func (req *CommentRequest) ValidateComment(c *gin.Context) bool {
	if err := c.Bind(&req); err != nil {
		errs := err.(validator.ValidationErrors)
		respErrors := make(map[string]interface{})
		for _, v := range errs {
			field, _ := reflect.TypeOf(&req).Elem().FieldByName(v.Field())
			fieldName, _ := field.Tag.Lookup("form")
			respErrors[fieldName] = v.Tag()
		}
		c.JSON(http.StatusBadRequest, gin.H{"errors": respErrors})
		return true
	}

	return false
}
