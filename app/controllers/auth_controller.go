package controllers

import (
	"go_bbs/form/api"
	"go_bbs/models/repository"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type AuthController struct{}

func (ac AuthController) Register(c *gin.Context) {
	req := api.RegisterRequest{}
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
		return
	}

	a := repository.AuthRepository{}
	a.CreateUser(c)

	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}
