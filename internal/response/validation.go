package response

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func WriteValidationError(c *gin.Context, err error) {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]FieldError, len(ve))
		for i, fe := range ve {
			out[i] = FieldError{
				Field:   toSnakeCase(fe.Field()),
				Message: validationMessage(fe),
			}
		}
		c.JSON(http.StatusBadRequest, gin.H{"errors": out})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}

func validationMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return toSnakeCase(fe.Field()) + " is required"
	case "min":
		return toSnakeCase(fe.Field()) + " must be at least " + fe.Param() + " characters"
	case "max":
		return toSnakeCase(fe.Field()) + " must be at most " + fe.Param() + " characters"
	case "eqfield":
		return toSnakeCase(fe.Field()) + " must be equal to " + toSnakeCase(fe.Param())
	default:
		return toSnakeCase(fe.Field()) + " is invalid"
	}
}

func toSnakeCase(str string) string {
	var sb strings.Builder
	for i, r := range str {
		if i > 0 && r >= 'A' && r <= 'Z' {
			sb.WriteRune('_')
		}
		sb.WriteRune(r)
	}
	return strings.ToLower(sb.String())
}
