package requests

import (
	"fmt"
	"go_bbs/repository"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RegisterRequest struct {
	Name     string `form:"name" binding:"required,max=255"`
	Password string `form:"password" binding:"required,min=6,max=13"`
	Token    string `form:"token" binding:"required"`
}

// Function to validate that it match the rule of PreRegisterRequest
func (req *RegisterRequest) ValidateRegister(c *gin.Context) bool {
	if err := c.Bind(&req); err != nil {
		errs := err.(validator.ValidationErrors)
		fmt.Println(errs)
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

func (req *RegisterRequest) ValidateEmailToken(c *gin.Context) bool {
	repo := repository.AuthRepository{}
	isExists, err := repo.IsPreUserExists(req.Token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": err})
		return true
	}

	if !isExists {
		respErrors := make(map[string]interface{})
		respErrors["token"] = "register link was expired"
		c.JSON(http.StatusBadRequest, gin.H{"errors": respErrors})
		return true
	}

	return false
}
