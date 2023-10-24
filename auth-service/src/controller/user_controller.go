package controller

import (
	"auth-service/src/exception"
	"auth-service/src/model/dto"
	"auth-service/src/service"
	"auth-service/src/service/service_interface"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserController struct {
	*BaseController
	service service_interface.UserServiceInterface
}

func NewUserController() *UserController {
	return &UserController{
		BaseController: NewBaseController(),
		service:        service.NewUserService(),
	}
}

func (this *UserController) Index(c echo.Context) error {
	users, err := this.service.GetAll()
	if err != nil {
		return err
	}

	response := make([]*dto.UserResponse, len(users))
	for ind, user := range users {
		response[ind] = dto.NewUserResponse(user)
	}

	return this.json(http.StatusOK, response, c)
}

func (this *UserController) Store(c echo.Context) error {
	request := dto.UserStoreRequest{}
	if err := this.handleRequest(&request, c); err != nil {
		return err
	}

	user, err := this.service.Create(&request)
	if err != nil {
		return err
	}

	return this.json(http.StatusOK, dto.NewUserResponse(user), c)
}

func (this *UserController) Show(c echo.Context) error {
	id, err := this.parseToUint(c.Param("id"), c)
	if err != nil {
		return err
	}

	authUser, err := this.authUser(c)
	if err != nil {
		return err
	}

	if authUser.ID != id {
		return exception.ErrNotPermission
	}

	user, err := this.service.GetById(id)
	if err != nil {
		return err
	}

	return this.json(http.StatusOK, dto.NewUserResponse(user), c)
}

func (this *UserController) UpdatePassword(c echo.Context) error {
	id, err := this.parseToUint(c.Param("id"), c)
	if err != nil {
		return err
	}

	authUser, err := this.authUser(c)
	if err != nil {
		return err
	}

	if authUser.ID != id {
		return exception.ErrNotPermission
	}

	request := dto.UserUpdatePasswordRequest{}
	if err = this.handleRequest(&request, c); err != nil {
		return err
	}

	user, err := this.service.UpdatePassword(id, &request)
	if err != nil {
		return err
	}

	return this.json(http.StatusOK, dto.NewUserResponse(user), c)
}

func (this *UserController) Delete(c echo.Context) error {
	id, err := this.parseToUint(c.Param("id"), c)
	if err != nil {
		return nil
	}

	authUser, err := this.authUser(c)
	if err != nil {
		return err
	}

	if authUser.ID != id {
		return exception.ErrNotPermission
	}

	if err = this.service.Delete(id); err != nil {
		return err
	}

	return this.json(http.StatusOK, nil, c)
}
