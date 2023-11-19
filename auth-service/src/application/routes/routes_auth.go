package routes

import (
	"auth-service/src/controller"
	"auth-service/src/middleware"
	"auth-service/src/service/service_i"
	"github.com/labstack/echo/v4"
)

func AuthRegisterRoute(
	group *echo.Group,
	authService service_i.AuthServiceI,
	authM *middleware.JwtMiddleware,
) {
	authController := controller.NewAuthController(authService)

	group.POST("/login", authController.Login)
	group.GET("/me", authController.Me, authM.Run())
}
