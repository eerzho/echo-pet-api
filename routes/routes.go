package routes

import (
	"echo-pet-api/src/controller"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
)

func RegisterRoutes(app *echo.Echo) {
	routesV1 := app.Group("/api/v1")
	jwtMiddleware := JWT([]byte("secret"))

	postController := &controller.PostController{}
	routesV1.GET("/posts", postController.Index)
	routesV1.POST("/posts", postController.Store, jwtMiddleware)
	routesV1.GET("/posts/:id", postController.Show)
	routesV1.PUT("/posts/:id", postController.Update, jwtMiddleware)
	routesV1.DELETE("/posts/:id", postController.Delete, jwtMiddleware)

	userController := &controller.UserController{}
	routesV1.GET("/users", userController.Index, jwtMiddleware)
	routesV1.POST("/users", userController.Store)
	routesV1.GET("/users/:id", userController.Show, jwtMiddleware)
	routesV1.PUT("/users/:id", userController.Update, jwtMiddleware)
	routesV1.DELETE("/users/:id", userController.Delete, jwtMiddleware)
	routesV1.POST("/login", userController.Login)
	routesV1.GET("/me", userController.Me, jwtMiddleware)
}

type jwtConfig struct {
	Skipper    skipper
	SigningKey interface{}
}

type skipper func(c echo.Context) bool

type jwtExtractor func(echo.Context) (string, error)

func JWT(key interface{}) echo.MiddlewareFunc {
	c := jwtConfig{}
	c.SigningKey = key
	return jwtWithConfig(c)
}

func jwtWithConfig(config jwtConfig) echo.MiddlewareFunc {
	extractor := jwtFromHeader("Authorization", "Bearer")
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			auth, err := extractor(c)
			if err != nil {
				if config.Skipper != nil {
					if config.Skipper(c) {
						return next(c)
					}
				}
				return c.JSON(http.StatusUnauthorized, err.Error())
			}
			token, err := jwt.Parse(auth, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return config.SigningKey, nil
			})
			if err != nil {
				return c.JSON(http.StatusForbidden, "Invalid or expired jwt")
			}
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				id := uint(claims["id"].(float64))
				email := claims["email"].(string)
				c.Set("userID", id)
				c.Set("userEmail", email)
				return next(c)
			}
			return c.JSON(http.StatusForbidden, "Invalid or expired jwt")
		}
	}
}

func jwtFromHeader(header string, authScheme string) jwtExtractor {
	return func(c echo.Context) (string, error) {
		auth := c.Request().Header.Get(header)
		l := len(authScheme)
		if len(auth) > l+1 && auth[:l] == authScheme {
			return auth[l+1:], nil
		}
		return "", echo.NewHTTPError(http.StatusUnauthorized, "Missing or malformed jwt")
	}
}
