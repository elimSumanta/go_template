package helper

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type errorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	case "alpha":
		return "Should not contain number or symbol"
	case "email":
		return "Should be in email format"
	case "alphanum":
		return "Invalid format"
	case "oneof":
		return "Unknown key on requested field"
	case "number":
		return "Page format is in number"
	}
	return "Unknown error"
}

func BodyValidator(schema interface{}, c *gin.Context) error {
	if err := c.ShouldBindJSON(&schema); err != nil {
		val(err, c)
		return err
	}
	return nil
}

func QueryValidator(schema interface{}, c *gin.Context) error {
	if err := c.ShouldBindQuery(schema); err != nil {
		val(err, c)
		return err
	}

	return nil
}

func val(err error, c *gin.Context) {
	var errorsData []errorMsg
	var val validator.ValidationErrors

	if errors.As(err, &val) {
		errorsData = make([]errorMsg, len(val))
		for i, fe := range val {
			errorsData[i] = errorMsg{fe.Field(), getErrorMsg(fe)}
		}
	}

	HttpErrorResponse(http.StatusBadRequest, BaseErrorResponseSchema{
		Code:    "INVALID_BODY",
		Message: "Missing Required Information",
		Errors:  errorsData,
	}, nil, c)
}
