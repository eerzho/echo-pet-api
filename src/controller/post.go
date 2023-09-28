package controller

import (
	"echo-pet-api/src/model/dto"
	"echo-pet-api/src/service"
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

	request := dto.PostStoreRequest{}
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if err := pc.validate.Struct(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	response, err := pc.service.Create(&request)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, response)
}

func (pc *PostController) Show(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	}

	response, err := pc.service.GetById(uint(id))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, response)
}

func (pc *PostController) ShowBySlug(c echo.Context) error {

	slug := c.Param("slug")

	response, err := pc.service.GetBySlug(slug)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, response)
}

func (pc *PostController) Update(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	}

	request := dto.PostUpdateRequest{}
	if err = c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if err = pc.validate.Struct(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	response, err := pc.service.Update(uint(id), &request)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, response)
}

func (pc *PostController) Delete(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	}

	err = pc.service.Delete(uint(id))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return echo.NewHTTPError(http.StatusOK, http.StatusText(http.StatusOK))
}
