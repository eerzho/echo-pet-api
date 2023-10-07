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
	"strconv"
)

type CommentController struct {
	service  *service.CommentService
	validate *validator.Validate
}

func NewCommentController() *CommentController {
	return &CommentController{
		service:  service.NewCommentService(),
		validate: validator.New(),
	}
}

func (cc *CommentController) Index(c echo.Context) error {
	response, err := cc.service.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, response)
}

func (cc *CommentController) Store(c echo.Context) error {
	authID, ok := c.Get("auth_id").(uint)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	request := dto.CommentStoreRequest{}
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if err := cc.validate.Struct(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	response, err := cc.service.Create(authID, &request)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, response)
}

func (cc *CommentController) Show(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	response, err := cc.service.GetById(uint(id))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound)
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, response)
}

func (cc *CommentController) Update(c echo.Context) error {
	authID, ok := c.Get("auth_id").(uint)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	request := dto.CommentUpdateRequest{}
	if err = c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if err = cc.validate.Struct(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	response, err := cc.service.Update(authID, uint(id), &request)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound)
	} else if errors.Is(err, &exception.PermissionDenied{}) {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, response)
}

func (cc *CommentController) Delete(c echo.Context) error {
	authID, ok := c.Get("auth_id").(uint)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	err = cc.service.Delete(authID, uint(id))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound)
	} else if errors.Is(err, &exception.PermissionDenied{}) {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, http.StatusText(http.StatusOK))
}
