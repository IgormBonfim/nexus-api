package middlewares

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/igormbonfim/nexus-api/internal/dtos/requests"
)

var routeValidationMap = map[string]interface{}{
	"/api/users": requests.CreateUserDto{},
}

func ValidatorMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		objType, exists := routeValidationMap[ctx.FullPath()]
		if !exists {
			ctx.Next()
			return
		}

		obj := reflect.New(reflect.TypeOf(objType)).Interface()

		err := ctx.ShouldBindJSON(obj)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON structure"})
			ctx.Abort()
			return
		}

		validate := validator.New()
		err = validate.Struct(obj)
		if err != nil {
			formattedErrors := formatValidationErrors(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"errors": formattedErrors})
			ctx.Abort()
			return
		}

		ctx.Set("validatedData", obj)

		ctx.Next()
	}
}

func formatValidationErrors(err error) map[string]string {
	formattedErrors := make(map[string]string)

	validationErrors, ok := err.(validator.ValidationErrors)

	if ok {
		for _, fieldErr := range validationErrors {
			formattedErrors[fieldErr.Field()] = fmt.Sprintf("Validation failed on '%s' rule", fieldErr.Tag())
		}
	}
	return formattedErrors
}
