package validation

import (
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func Validation(err error, req interface{}, c *gin.Context) map[string]interface{} {
	p(reflect.TypeOf(&req))

	return make(map[string]interface{})
	// request := req.(t)
	// errs := err.(validator.ValidationErrors)
	// respErrors := make(map[string]interface{})
	// for _, v := range errs {
	// 	field, _ := reflect.TypeOf(&req).Elem().FieldByName(v.Field())
	// 	fieldName, _ := field.Tag.Lookup("form")
	// 	respErrors[fieldName] = v.Tag()
	// }
	// return respErrors
}
func p(a ...interface{}) {
	fmt.Println(a...)
}
