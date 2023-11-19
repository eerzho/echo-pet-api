package routes

import (
	"auth-service/src/application"
	"auth-service/src/middleware"
	"auth-service/src/repository"
	"auth-service/src/service"
	"os"
)

func RegisterRoute(urlPrefix string) {
	group := application.GlobalApp.Group(urlPrefix)

	// repositories
	userRepository := repository.NewUserRepository()
	roleRepository := repository.NewRoleRepository()
	permissionRepository := repository.NewPermissionRepository()

	// services
	jwtService := service.NewJWTService()
	roleService := service.NewRoleService(roleRepository)
	userService := service.NewUserService(userRepository, roleService, jwtService)
	authService := service.NewAuthService(userService, jwtService)
	permissionService := service.NewPermissionService(permissionRepository)

	// middlewares
	authM := middleware.NewJwtMiddleware(userService, []byte(os.Getenv("JWT_SECRET")), "Bearer ")
	permissionM := middleware.NewPermissionMiddleware(userService)

	// routes
	AuthRegisterRoute(group, authService, authM)
	UserRegisterRoute(group, userService, authM, permissionM)
	RoleRegisterRoute(group, roleService, authM, permissionM)
	PermissionRegisterRoute(group, permissionService, authM, permissionM)
}
