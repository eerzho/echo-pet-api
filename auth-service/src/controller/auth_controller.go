package controller

import (
	"auth-service/src/dto"
	"auth-service/src/service/service_i"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthController struct {
	*BaseController
	authService service_i.AuthServiceI
}

func NewAuthController(authService service_i.AuthServiceI) *AuthController {
	return &AuthController{
		BaseController: NewBaseController(),
		authService:    authService,
	}
}

// Login
// @title Login
// @description Login a user
// @accept json
// @produce json
// @tags auth
// @param loginRequest body dto.LoginRequest true "Login Request"
// @success	200 {object} dto.JSONResult{data=dto.LoginResponse}
// @router /login [post]
func (this *AuthController) Login(c echo.Context) error {
	request := dto.LoginRequest{}
	if err := this.handleRequest(&request, c); err != nil {
		return err
	}

	token, err := this.authService.Login(&request)
	if err != nil {
		return err
	}

	return this.json(http.StatusOK, dto.NewLoginResponse(token), c)
}

// Me
// @title Me
// @description	Get user info
// @accept json
// @produce json
// @security ApiKeyAuth
// @tags auth
// @success 200	{object} dto.JSONResult{data=dto.UserResponse}
// @router /me [get]
func (this *AuthController) Me(c echo.Context) error {
	user, err := this.authUser(c)
	if err != nil {
		return err
	}

	user, err = this.authService.GetUserById(user.ID)
	if err != nil {
		return err
	}

	return this.json(http.StatusOK, dto.NewUserResponse(user), c)
}
