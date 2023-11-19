package routes

import (
	"auth-service/src/controller"
	"auth-service/src/middleware"
	"auth-service/src/service/service_i"
	"github.com/labstack/echo/v4"
)

func UserRegisterRoute(
	group *echo.Group,
	userService service_i.UserServiceI,
	authM *middleware.JwtMiddleware,
	permissionM *middleware.PermissionMiddleware,
) {
	userController := controller.NewUserController(userService)

	group.GET("/users", userController.Index, authM.Run(), permissionM.Run("users-index"))
	group.POST("/users", userController.Store)
	group.GET("/users/:id", userController.Show, authM.Run())
	group.PATCH("/users/:id", userController.UpdatePassword, authM.Run())
	group.DELETE("/users/:id", userController.Delete, authM.Run())
}
