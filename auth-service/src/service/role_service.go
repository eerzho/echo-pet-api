package service

import (
	"auth-service/src/dto"
	"auth-service/src/model"
	"auth-service/src/repository/repository_i"
	"github.com/gosimple/slug"
)

type RoleService struct {
	roleRepository repository_i.RoleRepositoryI
}

func NewRoleService(roleRepository repository_i.RoleRepositoryI) *RoleService {
	return &RoleService{roleRepository: roleRepository}
}

func (this *RoleService) GetAll() ([]*model.Role, error) {
	return this.roleRepository.GetAll()
}

func (this *RoleService) GetById(id uint) (*model.Role, error) {
	return this.roleRepository.GetById(id)
}

func (this *RoleService) GetBySlug(slug string) (*model.Role, error) {
	return this.roleRepository.GetBySlug(slug)
}

func (this *RoleService) Create(request *dto.RoleStoreRequest) (*model.Role, error) {
	role := model.Role{Name: request.Name, Slug: slug.Make(request.Name)}

	return this.roleRepository.Save(role)
}

func (this *RoleService) Update(id uint, request *dto.RoleUpdateRequest) (*model.Role, error) {
	role, err := this.roleRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	role.Name = request.Name
	role.Slug = slug.Make(request.Name)

	return this.roleRepository.Save(*role)
}

func (this *RoleService) Delete(id uint) error {
	return this.roleRepository.Delete(id)
}

func (this *RoleService) AddPermissions(id uint, request *dto.RoleAddPermissionsRequest) error {
	return this.roleRepository.AddPermissions(id, request.PermissionsID)
}

func (this *RoleService) RemovePermissions(id uint, request *dto.RoleRemovePermissionsRequest) error {
	return this.roleRepository.RemovePermissions(id, request.PermissionsID)
}
