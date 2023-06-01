package utilities

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ValidatorErrorResponse(ctx *gin.Context, StatusCode int, Method string, Error interface{}) {
	errResponse := ErrorResponse{
		StatusCode: StatusCode,
		Method:     Method,
		Errors:     Error,
	}

	ctx.JSON(StatusCode, errResponse)
	defer ctx.AbortWithStatus(StatusCode)
}

func ParseBindingError(bindingError error) interface{} {

	var validationErrors validator.ValidationErrors

	if errors.As(bindingError, &validationErrors) {

		output := make([]ErrorMsg, len(validationErrors))

		for index, fieldError := range validationErrors {
			output[index] = ErrorMsg{fieldError.Field(), getErrorMsg(fieldError)}
		}

		return output
	}

	return nil
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fe.Field() + " is required"
	case "email":
		return "Invalid email address"
	case "len":
		return fe.Field() + " should be of length " + fe.Param() + " digits"
	case "min":
		return fe.Field() + " should be at least of " + fe.Param() + " characters"
	case "max":
		return fe.Field() + " should not be more than " + fe.Param() + " characters"
	}
	return "Unknown error"
}
