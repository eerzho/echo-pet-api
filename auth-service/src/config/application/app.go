package application

import (
	"auth-service/src/config/exception"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

func CreateApp() *echo.Echo {
	e := echo.New()

	e.HTTPErrorHandler = errorHandler

	e.Logger.SetLevel(log.DEBUG)

	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.Logger())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	return e
}

func errorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError

	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		messages := make(map[string]string)
		for _, field := range validationErrors {
			fieldName := strings.ToLower(field.Field())
			messages[fieldName] = fmt.Sprintf("Field validation for '%s' failed on the '%s' tag", fieldName, field.Tag())
		}

		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": http.StatusText(http.StatusBadRequest),
			"data":    messages,
		})
	} else {
		var he *echo.HTTPError
		if errors.As(err, &he) {
			code = he.Code
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			code = http.StatusNotFound
		} else if errors.Is(err, exception.ErrUnauthorized) {
			code = http.StatusUnauthorized
		} else if errors.Is(err, exception.ErrInvalidLogin) ||
			errors.Is(err, exception.ErrInvalidParam) {
			code = http.StatusBadRequest
		} else if errors.Is(err, exception.ErrNotPermission) {
			code = http.StatusForbidden
		}

		c.JSON(code, map[string]interface{}{
			"message": err.Error(),
			"data":    nil,
		})
	}
}
