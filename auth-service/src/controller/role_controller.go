package controller

import (
	"auth-service/src/model/dto"
	"auth-service/src/service"
	"auth-service/src/service/service_interface"
	"github.com/labstack/echo/v4"
	"net/http"
)

type RoleController struct {
	*BaseController
	service service_interface.RoleServiceInterface
}

func NewRoleController() *RoleController {
	return &RoleController{
		BaseController: NewBaseController(),
		service:        service.NewRoleService(),
	}
}

func (this *RoleController) Index(c echo.Context) error {
	roles, err := this.service.GetAll()
	if err != nil {
		return err
	}

	response := make([]*dto.RoleResponse, len(roles))
	for ind, role := range roles {
		response[ind] = dto.NewRoleResponse(role)
	}

	return this.json(http.StatusOK, response, c)
}

func (this *RoleController) Store(c echo.Context) error {
	request := dto.RoleStoreRequest{}
	if err := this.handleRequest(&request, c); err != nil {
		return err
	}

	role, err := this.service.Create(&request)
	if err != nil {
		return err
	}

	return this.json(http.StatusOK, dto.NewRoleResponse(role), c)
}

func (this *RoleController) Show(c echo.Context) error {
	id, err := this.parseToUint(c.Param("id"), c)
	if err != nil {
		return err
	}

	role, err := this.service.GetById(id)
	if err != nil {
		return err
	}

	return this.json(http.StatusOK, dto.NewRoleResponse(role), c)
}

func (this *RoleController) Update(c echo.Context) error {
	id, err := this.parseToUint(c.Param("id"), c)
	if err != nil {
		return err
	}

	request := dto.RoleUpdateRequest{}
	if err = this.handleRequest(&request, c); err != nil {
		return err
	}

	role, err := this.service.Update(id, &request)
	if err != nil {
		return err
	}

	return this.json(http.StatusOK, dto.NewRoleResponse(role), c)
}

func (this *RoleController) Delete(c echo.Context) error {
	id, err := this.parseToUint(c.Param("id"), c)
	if err != nil {
		return err
	}

	if err = this.service.Delete(id); err != nil {
		return err
	}

	return this.json(http.StatusOK, nil, c)
}

func (this *RoleController) AddPermissions(c echo.Context) error {
	id, err := this.parseToUint(c.Param("id"), c)
	if err != nil {
		return err
	}

	request := dto.RoleAddPermissionsRequest{}
	if err = this.handleRequest(&request, c); err != nil {
		return err
	}

	if err = this.service.AddPermissions(id, &request); err != nil {
		return err
	}

	return this.json(http.StatusOK, nil, c)
}

func (this *RoleController) RemovePermissions(c echo.Context) error {
	id, err := this.parseToUint(c.Param("id"), c)
	if err != nil {
		return err
	}

	request := dto.RoleRemovePermissionsRequest{}
	if err = this.handleRequest(&request, c); err != nil {
		return err
	}

	if err = this.service.RemovePermissions(id, &request); err != nil {
		return err
	}

	return this.json(http.StatusOK, nil, c)
}
