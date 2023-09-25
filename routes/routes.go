package routes

import (
	"echo-pet-api/src/controller"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(app *echo.Echo) {
	routesV1 := app.Group("/api/v1")

	postController := &controller.PostController{}
	routesV1.GET("/posts", postController.Index)
	routesV1.POST("/posts", postController.Store)
	routesV1.GET("/posts/:id", postController.Show)
	routesV1.PUT("/posts/:id", postController.Update)
	routesV1.DELETE("/posts/:id", postController.Delete)

	userController := &controller.UserController{}
	routesV1.GET("/users", userController.Index)
	routesV1.POST("/users", userController.Store)
	routesV1.GET("/users/:id", userController.Show)
	routesV1.PUT("/users/:id", userController.Update)
	routesV1.DELETE("/users/:id", userController.Delete)
}
