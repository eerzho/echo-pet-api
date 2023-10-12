package controller

import (
	"blog-service/src/model/dto"
	"blog-service/src/service"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type UserController struct {
	service  *service.UserService
	validate *validator.Validate
}

func NewUserController() *UserController {
	return &UserController{
		service:  service.NewUserService(),
		validate: validator.New(),
	}
}

func (uc *UserController) Index(c echo.Context) error {

	response, err := uc.service.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, response)
}

func (uc *UserController) Store(c echo.Context) error {

	request := dto.UserStoreRequest{}
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if err := uc.validate.Struct(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	response, err := uc.service.Create(&request)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, response)
}

func (uc *UserController) Show(c echo.Context) error {

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	response, err := uc.service.GetById(uint(id))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound)
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, response)
}

func (uc *UserController) Update(c echo.Context) error {

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	request := dto.UserUpdateRequest{}
	if err = c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if err = uc.validate.Struct(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	response, err := uc.service.UpdatePassword(uint(id), &request)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound)
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, response)
}

func (uc *UserController) Delete(c echo.Context) error {

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	err = uc.service.Delete(uint(id))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound)
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, http.StatusText(http.StatusOK))
}
