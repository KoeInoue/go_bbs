package requests

import (
	"go_bbs/repository"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PreRegisterRequest struct {
	Email string `form:"email" binding:"required"`
}

// Function to validate that it match the rule of PreRegisterRequest
func (req *PreRegisterRequest) Validate(c *gin.Context) bool {
	if err := c.Bind(&req); err != nil {
		// TODO to commonalize validation
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

// Function to validate that email is not duplicated
func (req *PreRegisterRequest) ValidateDuplicateEmail(c *gin.Context) bool {
	repo := repository.AuthRepository{}
	email, err := repo.GetPreUserByEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": err})
		return true
	}

	if email != "" {
		respErrors := make(map[string]interface{})
		respErrors["email"] = "email is already in use"
		c.JSON(http.StatusBadRequest, gin.H{"errors": respErrors})
		return true
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": err})
		return true
	}

	return false
}
