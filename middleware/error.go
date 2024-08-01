package middleware

import (
	"fmt"
	"net/http"
	"notes/constant"
	"notes/helper"

	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Error() gin.HandlerFunc {
	return jsonErrorReporter(gin.ErrorTypeAny)
}

func jsonErrorReporter(errType gin.ErrorType) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		detectError := c.Errors.ByType(errType)
		if len(detectError) == 0 {
			return
		}

		err := detectError.Last()

		if handleUnauthorizedError(c, err) ||
			handleValidationErrors(c, err) ||
			handleNotFoundError(c, err) ||
			handleBadRequestError(c, err) ||
			handleServiceError(c, err) {
			return
		}
		handleInternalServerError(c)
	}
}

func handleUnauthorizedError(c *gin.Context, err error) bool {
	if errors.Is(err, constant.ErrUnAuth) {
		helper.SendResponse(c, nil, helper.Meta{
			StatusCode: http.StatusUnauthorized,
			Message:    "You're not authorized!",
		})
		return true
	}
	return false
}

func handleValidationErrors(c *gin.Context, err error) bool {
	if errs, ok := err.(validator.ValidationErrors); ok {
		errorMessages := make([]string, len(errs))
		for i, e := range errs {
			errorMessages[i] = fmt.Sprintf("Error in field %s, condition: %s", e.Field(), e.ActualTag())
		}
		helper.SendResponse(c, nil, helper.Meta{
			StatusCode: http.StatusBadRequest,
			Message:    errorMessages,
		})
		return true
	}
	return false
}

func handleNotFoundError(c *gin.Context, err error) bool {
	if errors.Is(err, constant.ErrNotFound) {
		errorMessage := helper.ErrMsgFormat(err)
		helper.SendResponse(c, nil, helper.Meta{
			StatusCode: http.StatusNotFound,
			Message:    errorMessage,
		})
		return true
	}
	return false
}

func handleBadRequestError(c *gin.Context, err error) bool {
	if errors.Is(err, constant.ErrBadRequest) {
		errorMessage := helper.ErrMsgFormat(err)
		helper.SendResponse(c, nil, helper.Meta{
			StatusCode: http.StatusBadRequest,
			Message:    errorMessage,
		})
		return true
	}
	return false
}

func handleServiceError(c *gin.Context, err error) bool {
	if errors.Is(err, constant.ErrService) {
		errorMessage := helper.ErrMsgFormat(err)
		helper.SendResponse(c, nil, helper.Meta{
			StatusCode: http.StatusBadRequest,
			Message:    errorMessage,
		})
		return true
	}
	return false
}

func handleInternalServerError(c *gin.Context) {
	helper.SendResponse(c, nil, helper.Meta{
		StatusCode: http.StatusInternalServerError,
		Message:    "INTERNAL SERVER ERROR",
	})
}
