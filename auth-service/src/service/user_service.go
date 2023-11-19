package service

import (
	"auth-service/src/dto"
	"auth-service/src/model"
	"auth-service/src/repository/repository_i"
	"auth-service/src/service/service_i"
)

type UserService struct {
	userRepository repository_i.UserRepositoryI
	roleService    service_i.RoleServiceI
	jwtService     service_i.JWTServiceI
}

func NewUserService(
	userRepository repository_i.UserRepositoryI,
	roleService service_i.RoleServiceI,
	jwtService service_i.JWTServiceI,
) *UserService {
	return &UserService{
		userRepository: userRepository,
		roleService:    roleService,
		jwtService:     jwtService,
	}
}

func (this *UserService) GetAll() ([]*model.User, error) {
	return this.userRepository.GetAll()
}

func (this *UserService) GetById(id uint) (*model.User, error) {
	return this.userRepository.GetById(id)
}

func (this *UserService) GetByEmail(email string) (*model.User, error) {
	return this.userRepository.GetByEmail(email)
}

func (this *UserService) Create(request *dto.UserStoreRequest) (*model.User, error) {
	passwordHash, err := this.jwtService.GenerateHash(request.Password)
	if err != nil {
		return nil, err
	}

	role, err := this.roleService.GetBySlug("user")
	if err != nil {
		return nil, err
	}

	user := model.User{Email: request.Email, Name: request.Name, Password: passwordHash, RoleID: role.ID}

	return this.userRepository.Save(user)
}

func (this *UserService) UpdatePassword(id uint, request *dto.UserUpdatePasswordRequest) (*model.User, error) {
	user, err := this.userRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	user.Password, err = this.jwtService.GenerateHash(request.Password)
	if err != nil {
		return nil, err
	}

	return this.userRepository.Save(*user)
}

func (this *UserService) Delete(id uint) error {
	return this.userRepository.Delete(id)
}

func (this *UserService) HasPermission(id uint, permissionSlug string) bool {
	return this.userRepository.HasPermission(id, permissionSlug)
}
