package service

import (
	"auth-service/src/dto"
	"auth-service/src/model"
	"auth-service/src/repository/repository_i"
	"github.com/gosimple/slug"
)

type PermissionService struct {
	permissionRepository repository_i.PermissionRepositoryI
}

func NewPermissionService(permissionRepository repository_i.PermissionRepositoryI) *PermissionService {
	return &PermissionService{permissionRepository: permissionRepository}
}

func (this *PermissionService) GetAllByRole(roleID uint) ([]*model.Permission, error) {
	return this.permissionRepository.GetAllByRole(roleID)
}

func (this *PermissionService) GetById(id uint) (*model.Permission, error) {
	return this.permissionRepository.GetById(id)
}

func (this *PermissionService) Create(request *dto.PermissionStoreRequest) (*model.Permission, error) {
	permission := model.Permission{Name: request.Name, Slug: slug.Make(request.Name)}

	return this.permissionRepository.Save(permission)
}

func (this *PermissionService) Update(id uint, request *dto.PermissionUpdateRequest) (*model.Permission, error) {
	permission, err := this.permissionRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	permission.Name = request.Name
	permission.Slug = slug.Make(request.Name)

	return this.permissionRepository.Save(*permission)
}

func (this *PermissionService) Delete(id uint) error {
	return this.permissionRepository.Delete(id)
}
