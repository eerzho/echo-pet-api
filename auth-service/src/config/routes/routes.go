package routes

import (
	"auth-service/src/config/routes/middleware"
	"auth-service/src/controller"
	"github.com/labstack/echo/v4"
	"os"
)

func RegisterRoutes(app *echo.Echo) {
	routesV1 := app.Group("/api/v1")
	authM := middleware.NewJwtMiddleware([]byte(os.Getenv("JWT_SECRET")), "Bearer ")
	permissionM := middleware.NewPermissionMiddleware()

	authController := controller.NewAuthController()
	routesV1.POST("/login", authController.Login)
	routesV1.GET("/me", authController.Me, authM.Run())

	userController := controller.NewUserController()
	routesV1.GET("/users", userController.Index, authM.Run(), permissionM.Run("users-index"))
	routesV1.POST("/users", userController.Store)
	routesV1.GET("/users/:id", userController.Show, authM.Run())
	routesV1.PATCH("/users/:id", userController.UpdatePassword, authM.Run())
	routesV1.DELETE("/users/:id", userController.Delete, authM.Run())

	roleController := controller.NewRoleController()
	routesV1.GET("/roles", roleController.Index, authM.Run(), permissionM.Run("roles-index"))
	routesV1.POST("/roles", roleController.Store, authM.Run(), permissionM.Run("roles-store"))
	routesV1.GET("/roles/:id", roleController.Show, authM.Run(), permissionM.Run("roles-show"))
	routesV1.PATCH("/roles/:id", roleController.Update, authM.Run(), permissionM.Run("roles-update"))
	routesV1.DELETE("/roles/:id", roleController.Delete, authM.Run(), permissionM.Run("roles-delete"))
	routesV1.POST("/roles/:id/permissions", roleController.AddPermissions, authM.Run(), permissionM.Run("roles-add-permissions"))
	routesV1.DELETE("/roles/:id/permissions", roleController.RemovePermissions, authM.Run(), permissionM.Run("roles-remove-permissions"))

	permissionController := controller.NewPermissionController()
	routesV1.GET("/permissions", permissionController.Index, authM.Run(), permissionM.Run("permissions-index"))
	routesV1.POST("/permissions", permissionController.Store, authM.Run(), permissionM.Run("permissions-store"))
	routesV1.GET("/permissions/:id", permissionController.Show, authM.Run(), permissionM.Run("permissions-show"))
	routesV1.PATCH("/permissions/:id", permissionController.Update, authM.Run(), permissionM.Run("permissions-update"))
	routesV1.DELETE("/permissions/:id", permissionController.Delete, authM.Run(), permissionM.Run("permissions-delete"))
}
