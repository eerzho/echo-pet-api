package middleware

import (
	"auth-service/src/exception"
	"auth-service/src/model"
	"auth-service/src/service"
	"auth-service/src/service/service_interface"
	"github.com/labstack/echo/v4"
)

type PermissionMiddleware struct {
	userService service_interface.UserServiceInterface
}

func NewPermissionMiddleware() *PermissionMiddleware {
	return &PermissionMiddleware{userService: service.NewUserService()}
}

func (this *PermissionMiddleware) Run(permissionSlug string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user, ok := c.Get("auth_user").(*model.User)
			if !ok {
				return exception.ErrUnauthorized
			}

			if !this.userService.HasPermission(user.ID, permissionSlug) {
				return exception.ErrNotPermission
			}

			return next(c)
		}
	}
}
