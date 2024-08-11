package controller

import (
	"net/http"
	"notes/internal/dtos"
	"notes/internal/helper"
	"notes/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	LoginWithGoogle(c *gin.Context)
}

type AuthControllerImpl struct {
	AuthService service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return &AuthControllerImpl{
		AuthService: authService,
	}
}

func (cl *AuthControllerImpl) LoginWithGoogle(c *gin.Context) {
	req := dtos.GoogleLoginReq{}

	err := c.Bind(&req)
	if err != nil {
		c.Error(err)
		return
	}

	res, err := cl.AuthService.GoogleOAuth(c.Request.Context(), req)
	if err != nil {
		c.Error(err)
		return
	}
	helper.SendResponse(c, res, helper.Meta{
		StatusCode: http.StatusOK,
		Message:    "Sign in with google succesfully",
	})

}
