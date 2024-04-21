package handler

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func TemplateResponse(c *gin.Context, code int, name string, obj gin.H) {
	if obj == nil {
		obj = gin.H{}
	}
	obj["LIVE_RELOAD"] = os.Getenv("LIVE_RELOAD")
	c.HTML(code, name, obj)
}

// ShouldValidate validates data using go-playground validator
// and sends back user friendly error
func ShouldValidate(c *gin.Context, obj any) map[string]string {
	validate := validator.New()

	// Use JSON tag as the field name in validation errors
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("form"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	err := validate.Struct(obj)
	if err != nil {
		log.Printf("ShouldValidate error: %+v", err)
		return UserFriendlyError(err)
	}
	return nil
}

func UserFriendlyError(err error) map[string]string {
	errs := map[string]string{}

	for _, err := range err.(validator.ValidationErrors) {
		fieldName := err.Field()

		switch err.Tag() {
		case "required":
			errs[fieldName] = "This field is required."
		case "email":
			errs[fieldName] = "This field must be a valid email address."
		case "min":
			errs[fieldName] = fmt.Sprintf("This field must be at least %s characters long.", err.Param())
		case "max":
			errs[fieldName] = fmt.Sprintf("This field must not exceed %s characters.", err.Param())
		default:
			errs[fieldName] = "This field is not valid."
		}
	}

	return errs
}
