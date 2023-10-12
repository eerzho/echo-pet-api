package controller

import (
	"blog-service/src/exception"
	"blog-service/src/model/dto"
	"blog-service/src/service"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type PostController struct {
	service  *service.PostService
	validate *validator.Validate
}

func NewPostController() *PostController {
	return &PostController{
		service:  service.NewPostService(),
		validate: validator.New(),
	}
}

func (pc *PostController) Index(c echo.Context) error {

	response, err := pc.service.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, response)
}

func (pc *PostController) Store(c echo.Context) error {
	authorID, ok := c.Get("auth_id").(uint)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	request := dto.PostStoreRequest{}
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if err := pc.validate.Struct(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	response, err := pc.service.Create(authorID, &request)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, response)
}

func (pc *PostController) Show(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	response, err := pc.service.GetById(uint(id))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound)
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, response)
}

func (pc *PostController) ShowBySlug(c echo.Context) error {

	slug := c.Param("slug")

	response, err := pc.service.GetBySlug(slug)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound)
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, response)
}

func (pc *PostController) Update(c echo.Context) error {
	authorID, ok := c.Get("auth_id").(uint)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	request := dto.PostUpdateRequest{}
	if err = c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if err = pc.validate.Struct(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	response, err := pc.service.Update(authorID, uint(id), &request)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound)
	} else if errors.Is(err, &exception.PermissionDenied{}) {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, response)
}

func (pc *PostController) Delete(c echo.Context) error {
	authorID, ok := c.Get("auth_id").(uint)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	err = pc.service.Delete(authorID, uint(id))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound)
	} else if errors.Is(err, &exception.PermissionDenied{}) {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, http.StatusText(http.StatusOK))
}
