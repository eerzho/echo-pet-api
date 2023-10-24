package controller

import (
	"auth-service/src/model/dto"
	"auth-service/src/service"
	"auth-service/src/service/service_interface"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthController struct {
	*BaseController
	service service_interface.AuthServiceInterface
}

func NewAuthController() *AuthController {
	return &AuthController{
		BaseController: NewBaseController(),
		service:        service.NewAuthService(),
	}
}

func (this *AuthController) Login(c echo.Context) error {
	request := dto.LoginRequest{}
	if err := this.handleRequest(&request, c); err != nil {
		return err
	}

	token, err := this.service.Login(&request)
	if err != nil {
		return err
	}

	return this.json(http.StatusOK, dto.NewLoginResponse(token), c)
}

func (this *AuthController) Me(c echo.Context) error {
	user, err := this.authUser(c)
	if err != nil {
		return err
	}

	return this.json(http.StatusOK, dto.NewUserResponse(user), c)
}
