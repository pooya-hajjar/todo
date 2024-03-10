package apiErrors

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pooya-hajjar/todo/constants/validations"
	"net/http"
)

type ValidationErrors struct {
	Field   string
	Message string
}

type ServerErrors struct {
	Message string
}

type QueryErrors struct {
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
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ValidationErrors, len(ve))

			for i, fe := range ve {
				out[i] = ValidationErrors{fe.Field(), ErrorMessageForTag(fe.Tag(), fe.Param())}
			}
			ctx.JSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
	}
}
