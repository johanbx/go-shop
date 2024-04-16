package handler

import (
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// ShouldPayloadBind binds data from the HTTP request to obj using gin ShouldBind
// obj must be a pointer to a struct for binding to work correctly.
func ShouldPayloadBind(c *gin.Context, obj any) error {
	log.Printf("ShouldPayloadBind %+v", obj)
	var err error
	if err = c.ShouldBind(obj); err != nil {
		log.Printf("Request error: %+v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
	}
	return err
}

// ShouldValidate validates data using go-playground validator
// and sends back user friendly error
func ShouldValidate(c *gin.Context, obj any) error {
	log.Printf("ShouldValidate %+v", obj)
	var err error
	validate := validator.New()
	if err = validate.Struct(obj); err != nil {
		log.Printf("Validation error: %+v", err)
		c.JSON(http.StatusBadRequest, UserFriendlyError(err))
	}
	return err
}

func UserFriendlyError(err error) map[string]string {
	errs := map[string]string{}

	for _, err := range err.(validator.ValidationErrors) {
		// Use json tag as field name if available, otherwise use the struct's field name
		fieldName := err.Field()
		formFieldName := fieldName // Fallback to the actual struct field name

		// Check if a json tag is available for the field
		if fieldSpecs, ok := reflect.TypeOf(CatalogItemRequestModel{}).FieldByName(fieldName); ok {
			// Get the json tag value, if it's "-", we will skip it
			if tag := fieldSpecs.Tag.Get("form"); tag != "-" && tag != "" {
				formFieldName = tag
			}
		}

		// You can add more cases to cover different validation scenarios
		switch err.Tag() {
		case "required":
			errs[formFieldName] = "This field is required."
		case "email":
			errs[formFieldName] = "This field must be a valid email address."
		case "min":
			errs[formFieldName] = fmt.Sprintf("This field must be at least %s characters long.", err.Param())
		case "max":
			errs[formFieldName] = fmt.Sprintf("This field must not exceed %s characters.", err.Param())
		// Add other validations as necessary
		default:
			// Generic default message
			errs[formFieldName] = "This field is not valid."
		}
	}

	return errs
}
