package apiErrors

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pooya-hajjar/todo/constants/validations"
)

type ValidationErrors struct {
	Field   string
	Message string
}

func ErrorMessageForTag(tag string, value string) string {
	switch tag {
	case "required":
		return validations.Required

	case "min":
		return fmt.Sprintf(validations.Min, value)
	case "max":
		return fmt.Sprintf(validations.Max, value)

	case "email":
		return validations.Email

	default:
		return validations.Default
	}
}

func HandleValidationError(ctx *gin.Context, err error) {
	var ve validator.ValidationErrors

	if errors.As(err, &ve) {
		out := make(map[string]string)

		for _, fe := range ve {

			out[fe.Field()] = ErrorMessageForTag(fe.Tag(), fe.Param())
		}

		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": out})
		return
	}

	ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "empty input"})
}
