package routes

import (
	"auth-service/src/config/routes/middleware"
	"auth-service/src/controller"
	"github.com/labstack/echo/v4"
	"os"
)

func RegisterRoutes(app *echo.Echo) {
	routesV1 := app.Group("/api/v1")
	authMiddleware := middleware.NewJwtMiddleware([]byte(os.Getenv("JWT_SECRET")), "Bearer ")

	authController := controller.NewAuthController()
	routesV1.POST("/login", authController.Login)
	routesV1.GET("/me", authController.Me, authMiddleware.Run())

	userController := controller.NewUserController()
	routesV1.GET("/users", userController.Index)
	routesV1.POST("/users", userController.Store)
	routesV1.GET("/users/:id", userController.Show)
	routesV1.PATCH("/users/:id", userController.UpdatePassword)
	routesV1.DELETE("/users/:id", userController.Delete)
}
