package controller

import (
	"echo-pet-api/src/exception"
	"echo-pet-api/src/model/dto"
	"echo-pet-api/src/service"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type AuthController struct {
	service     *service.AuthService
	userService *service.UserService
	validate    *validator.Validate
}

func NewAuthController() *AuthController {
	return &AuthController{
		service:     service.NewAuthService(),
		userService: service.NewUserService(),
		validate:    validator.New(),
	}
}

func (ac *AuthController) Login(c echo.Context) error {
	request := dto.LoginRequest{}
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if err := ac.validate.Struct(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	response, err := ac.service.Login(&request)
	if errors.Is(err, &exception.InvalidLoginError{}) {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, response)
}

func (ac *AuthController) Me(c echo.Context) error {
	id, ok := c.Get("auth_id").(uint)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	response, err := ac.userService.GetById(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound)
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, response)
}
