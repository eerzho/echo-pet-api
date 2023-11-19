package controller

import (
	"auth-service/src/dto"
	"auth-service/src/service/service_i"
	"github.com/labstack/echo/v4"
	"net/http"
)

type RoleController struct {
	*BaseController
	roleService service_i.RoleServiceI
}

func NewRoleController(roleService service_i.RoleServiceI) *RoleController {
	return &RoleController{
		BaseController: NewBaseController(),
		roleService:    roleService,
	}
}

// Index
// @title Index
// @description List of roles
// @accept json
// @produce json
// @security ApiKeyAuth
// @tags role
// @success 200 {object} dto.JSONResult{data=[]dto.RoleResponse}
// @router /roles [get]
func (this *RoleController) Index(c echo.Context) error {
	roles, err := this.roleService.GetAll()
	if err != nil {
		return err
	}

	response := make([]*dto.RoleResponse, len(roles))
	for ind, role := range roles {
		response[ind] = dto.NewRoleResponse(role)
	}

	return this.json(http.StatusOK, response, c)
}

// Store
// @title Store
// @description Create a role
// @accept json
// @produce json
// @security ApiKeyAuth
// @tags role
// @param roleStoreRequest body dto.RoleStoreRequest true "Role store request"
// @success 200 {object} dto.JSONResult{data=dto.RoleResponse}
// @router /roles [post]
func (this *RoleController) Store(c echo.Context) error {
	request := dto.RoleStoreRequest{}
	if err := this.handleRequest(&request, c); err != nil {
		return err
	}

	role, err := this.roleService.Create(&request)
	if err != nil {
		return err
	}

	return this.json(http.StatusOK, dto.NewRoleResponse(role), c)
}

// Show
// @title Show
// @description Get a role
// @accept json
// @produce json
// @security ApiKeyAuth
// @tags role
// @param id path int true "Role ID"
// @success 200 {object} dto.JSONResult{data=dto.RoleResponse}
// @router /roles/{id} [get]
func (this *RoleController) Show(c echo.Context) error {
	id, err := this.parseToUint(c.Param("id"))
	if err != nil {
		return err
	}

	role, err := this.roleService.GetById(id)
	if err != nil {
		return err
	}

	return this.json(http.StatusOK, dto.NewRoleResponse(role), c)
}

// Update
// @title Update
// @description Update a role
// @accept json
// @produce json
// @security ApiKeyAuth
// @tags role
// @param roleUpdateRequest body dto.RoleUpdateRequest true "Role update request"
// @param id path int true "Role ID"
// @success 200 {object} dto.JSONResult{data=dto.PermissionResponse}
// @router /roles/{id} [patch]
func (this *RoleController) Update(c echo.Context) error {
	id, err := this.parseToUint(c.Param("id"))
	if err != nil {
		return err
	}

	request := dto.RoleUpdateRequest{}
	if err = this.handleRequest(&request, c); err != nil {
		return err
	}

	role, err := this.roleService.Update(id, &request)
	if err != nil {
		return err
	}

	return this.json(http.StatusOK, dto.NewRoleResponse(role), c)
}

// Delete
// @title Delete
// @description Delete a role
// @accept json
// @produce json
// @security ApiKeyAuth
// @tags role
// @param id path int true "Role ID"
// @success 200 {object} dto.JSONResult
// @router /roles/{id} [delete]
func (this *RoleController) Delete(c echo.Context) error {
	id, err := this.parseToUint(c.Param("id"))
	if err != nil {
		return err
	}

	if err = this.roleService.Delete(id); err != nil {
		return err
	}

	return this.json(http.StatusOK, nil, c)
}

// AddPermissions
// @title AddPermissions
// @description Add permissions to role
// @accept json
// @produce json
// @security ApiKeyAuth
// @tags role
// @param id path int true "Role ID"
// @param roleAddPermissionsRequest body dto.RoleAddPermissionsRequest true "Role add permissions request"
// @success 200 {object} dto.JSONResult
// @router /role/{id}/permissions [post]
func (this *RoleController) AddPermissions(c echo.Context) error {
	id, err := this.parseToUint(c.Param("id"))
	if err != nil {
		return err
	}

	request := dto.RoleAddPermissionsRequest{}
	if err = this.handleRequest(&request, c); err != nil {
		return err
	}

	if err = this.roleService.AddPermissions(id, &request); err != nil {
		return err
	}

	return this.json(http.StatusOK, nil, c)
}

// RemovePermissions
// @title RemovePermissions
// @description Remove permissions from role
// @accept json
// @produce json
// @security ApiKeyAuth
// @tags role
// @param id path int true "Role ID"
// @param roleRemovePermissionsRequest body dto.RoleRemovePermissionsRequest true "Role remove permissions request"
// @success 200 {object} dto.JSONResult
// @router /role/{id}/permissions [delete]
func (this *RoleController) RemovePermissions(c echo.Context) error {
	id, err := this.parseToUint(c.Param("id"))
	if err != nil {
		return err
	}

	request := dto.RoleRemovePermissionsRequest{}
	if err = this.handleRequest(&request, c); err != nil {
		return err
	}

	if err = this.roleService.RemovePermissions(id, &request); err != nil {
		return err
	}

	return this.json(http.StatusOK, nil, c)
}
