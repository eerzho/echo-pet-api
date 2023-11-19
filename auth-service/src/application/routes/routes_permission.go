package routes

import (
	"auth-service/src/controller"
	"auth-service/src/middleware"
	"auth-service/src/service/service_i"
	"github.com/labstack/echo/v4"
)

func PermissionRegisterRoute(
	group *echo.Group,
	permissionService service_i.PermissionServiceI,
	authM *middleware.JwtMiddleware,
	permissionM *middleware.PermissionMiddleware,
) {
	permissionController := controller.NewPermissionController(permissionService)

	group.GET("/permissions", permissionController.Index, authM.Run(), permissionM.Run("permissions-index"))
	group.POST("/permissions", permissionController.Store, authM.Run(), permissionM.Run("permissions-store"))
	group.GET("/permissions/:id", permissionController.Show, authM.Run(), permissionM.Run("permissions-show"))
	group.PATCH("/permissions/:id", permissionController.Update, authM.Run(), permissionM.Run("permissions-update"))
	group.DELETE("/permissions/:id", permissionController.Delete, authM.Run(), permissionM.Run("permissions-delete"))
}
