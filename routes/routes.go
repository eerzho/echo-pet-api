package routes

import (
	"echo-pet-api/routes/middleware"
	"echo-pet-api/src/controller"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(app *echo.Echo) {
	routesV1 := app.Group("/api/v1")

	postController := controller.NewPostController()
	routesV1.GET("/posts", postController.Index)
	routesV1.POST("/posts", postController.Store, middleware.JWTMiddleware())
	routesV1.GET("/posts/:id", postController.Show)
	routesV1.GET("/posts/slug/:slug", postController.ShowBySlug)
	routesV1.PUT("/posts/:id", postController.Update, middleware.JWTMiddleware())
	routesV1.DELETE("/posts/:id", postController.Delete, middleware.JWTMiddleware())

	userController := controller.NewUserController()
	routesV1.GET("/users", userController.Index, middleware.JWTMiddleware())
	routesV1.POST("/users", userController.Store)
	routesV1.GET("/users/:id", userController.Show, middleware.JWTMiddleware())
	routesV1.PUT("/users/:id", userController.Update, middleware.JWTMiddleware())
	routesV1.DELETE("/users/:id", userController.Delete, middleware.JWTMiddleware())

	authController := controller.NewAuthController()
	routesV1.POST("/login", authController.Login)
	routesV1.GET("/me", authController.Me, middleware.JWTMiddleware())
}
