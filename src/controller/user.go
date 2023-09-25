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

type UserController struct {
}

func (uc *UserController) Index(c echo.Context) error {
	var users []model.User

	db := database.Connection()
	if err := db.Find(&users).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Server error")
	}

	var userList []*model.UserResponse
	for _, user := range users {
		userList = append(userList, model.NewUserResponseFromModel(&user))
	}

	return c.JSON(http.StatusOK, userList)
}

func (uc *UserController) Store(c echo.Context) error {
	req := model.UserStoreRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusInternalServerError, "Server error")
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	user := model.NewUserFromStoreRequest(&req)

	db := database.Connection()
	if err := db.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Server error")
	}

	return c.JSON(http.StatusOK, model.NewUserResponseFromModel(user))
}

func (uc *UserController) Show(c echo.Context) error {
	var id int
	if result, err := strconv.Atoi(c.Param("id")); err != nil {
		return c.JSON(http.StatusNotFound, "Not found")
	} else {
		id = result
	}

	var user model.User

	db := database.Connection()
	if err := db.First(&user, id).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, "Not found")
	}

	return c.JSON(http.StatusOK, model.NewUserResponseFromModel(&user))
}

func (uc *UserController) Update(c echo.Context) error {
	var id int
	if result, err := strconv.Atoi(c.Param("id")); err != nil {
		return c.JSON(http.StatusNotFound, "Not found")
	} else {
		id = result
	}

	req := model.UserUpdateRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusInternalServerError, "Server error")
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var user model.User

	db := database.Connection()
	if err := db.First(&user, id).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, "Not found")
	}

	user.Password = req.Password

	if err := db.Save(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Server error")
	}

	return c.JSON(http.StatusOK, model.NewUserResponseFromModel(&user))
}

func (uc *UserController) Delete(c echo.Context) error {
	var id int
	if result, err := strconv.Atoi(c.Param("id")); err != nil {
		return c.JSON(http.StatusNotFound, "Not found")
	} else {
		id = result
	}

	var user model.User

	db := database.Connection()
	if err := db.First(&user, id).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, "Not found")
	}

	if err := db.Delete(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Server error")
	}

	return c.JSON(http.StatusOK, "Deleted")
}
