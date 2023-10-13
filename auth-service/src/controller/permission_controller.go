package controller

import (
	"auth-service/src/model/dto"
	"auth-service/src/service"
	"github.com/davecgh/go-spew/spew"
	"github.com/labstack/echo/v4"
	"net/http"
)

type PermissionController struct {
	*BaseController
	service *service.PermissionService
}

func NewPermissionController() *PermissionController {
	return &PermissionController{
		BaseController: NewBaseController(),
		service:        service.NewPermissionService(),
	}
}

func (this *PermissionController) Index(c echo.Context) error {
	roleID, err := this.parseToUint(c.QueryParam("role_id"), c)
	if err != nil {
		return err
	}

	spew.Dump(roleID)

	permissions, err := this.service.GetAllByRole(roleID)
	if err != nil {
		return err
	}

	response := make([]*dto.PermissionResponse, len(permissions))
	for ind, permission := range permissions {
		response[ind] = dto.NewPermissionResponse(permission)
	}

	return this.json(http.StatusOK, response, c)
}

func (this *PermissionController) Store(c echo.Context) error {
	request := dto.PermissionStoreRequest{}
	if err := this.handleRequest(&request, c); err != nil {
		return err
	}

	permission, err := this.service.Create(&request)
	if err != nil {
		return err
	}

	return this.json(http.StatusOK, dto.NewPermissionResponse(permission), c)
}

func (this *PermissionController) Show(c echo.Context) error {
	id, err := this.parseToUint(c.Param("id"), c)
	if err != nil {
		return err
	}

	permission, err := this.service.GetById(id)
	if err != nil {
		return err
	}

	return this.json(http.StatusOK, dto.NewPermissionResponse(permission), c)
}

func (this *PermissionController) Update(c echo.Context) error {
	id, err := this.parseToUint(c.Param("id"), c)
	if err != nil {
		return err
	}

	request := dto.PermissionUpdateRequest{}
	if err = this.handleRequest(&request, c); err != nil {
		return err
	}

	permission, err := this.service.Update(id, &request)
	if err != nil {
		return err
	}

	return this.json(http.StatusOK, dto.NewPermissionResponse(permission), c)
}

func (this *PermissionController) Delete(c echo.Context) error {
	id, err := this.parseToUint(c.Param("id"), c)
	if err != nil {
		return err
	}

	if err = this.service.Delete(id); err != nil {
		return err
	}

	return this.json(http.StatusOK, nil, c)
}
