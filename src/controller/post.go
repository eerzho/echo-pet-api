package controller

import (
	"echo-pet-api/database"
	"echo-pet-api/src/model"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type PostController struct {
}

func (pc *PostController) Index(c echo.Context) error {
	var posts []model.Post

	db := database.Connection()
	if err := db.Find(&posts).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Server error")
	}

	var postList []*model.PostResponse
	for _, post := range posts {
		postList = append(postList, model.NewPostResponseFromModel(&post))
	}

	return c.JSON(http.StatusOK, postList)
}

func (pc *PostController) Store(c echo.Context) error {

	req := model.PostStoreRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusInternalServerError, "Server error")
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	post := model.NewPostFromStoreRequest(&req)

	db := database.Connection()
	if err := db.Create(&post).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Server error")
	}

	return c.JSON(http.StatusOK, model.NewPostResponseFromModel(post))
}

func (pc *PostController) Show(c echo.Context) error {
	var id int
	if result, err := strconv.Atoi(c.Param("id")); err != nil {
		return c.JSON(http.StatusNotFound, "Not found")
	} else {
		id = result
	}

	var post model.Post

	db := database.Connection()
	if err := db.First(&post, id).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, "Not found")
	}

	return c.JSON(http.StatusOK, model.NewPostResponseFromModel(&post))
}

func (pc *PostController) Update(c echo.Context) error {
	var id int
	if result, err := strconv.Atoi(c.Param("id")); err != nil {
		return c.JSON(http.StatusNotFound, "Not found")
	} else {
		id = result
	}

	req := model.PostUpdateRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusInternalServerError, "Server error")
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var post model.Post

	db := database.Connection()
	if err := db.First(&post, id).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, "Server error")
	}

	post.Desc = req.Desc

	if err := db.Save(&post).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Server error")
	}

	return c.JSON(http.StatusOK, model.NewPostResponseFromModel(&post))
}

func (pc *PostController) Delete(c echo.Context) error {
	var id int
	if result, err := strconv.Atoi(c.Param("id")); err != nil {
		return c.JSON(http.StatusNotFound, "Not found")
	} else {
		id = result
	}

	var post model.Post

	db := database.Connection()
	if err := db.First(&post, id).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, "Not found")
	}

	if err := db.Delete(&post).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Server error")
	}

	return c.JSON(http.StatusOK, "Deleted")
}
