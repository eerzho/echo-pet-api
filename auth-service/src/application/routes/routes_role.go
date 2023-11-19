package routes

import (
	"auth-service/src/controller"
	"auth-service/src/middleware"
	"auth-service/src/service/service_i"
	"github.com/labstack/echo/v4"
)

func RoleRegisterRoute(
	group *echo.Group,
	roleService service_i.RoleServiceI,
	authM *middleware.JwtMiddleware,
	permissionM *middleware.PermissionMiddleware,
) {
	roleController := controller.NewRoleController(roleService)

	group.GET("/roles", roleController.Index, authM.Run(), permissionM.Run("roles-index"))
	group.POST("/roles", roleController.Store, authM.Run(), permissionM.Run("roles-store"))
	group.GET("/roles/:id", roleController.Show, authM.Run(), permissionM.Run("roles-show"))
	group.PATCH("/roles/:id", roleController.Update, authM.Run(), permissionM.Run("roles-update"))
	group.DELETE("/roles/:id", roleController.Delete, authM.Run(), permissionM.Run("roles-delete"))
	group.POST("/roles/:id/permissions", roleController.AddPermissions, authM.Run(), permissionM.Run("roles-add-permissions"))
	group.DELETE("/roles/:id/permissions", roleController.RemovePermissions, authM.Run(), permissionM.Run("roles-remove-permissions"))
}
