package service

import (
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/model"
	"svc-monitoring-maintenance/repository"
)

type RoleServiceInterface interface {
	GetRole() ([]entity.Role, error)
	CreateRole(request model.CreateRoleRequest) (entity.Role, error)
}

type roleService struct {
	repository repository.RoleRepositoryInteface
}

func RoleService(repository repository.RoleRepositoryInteface) *roleService {
	return &roleService{repository}
}

func (roleService *roleService) GetRole() ([]entity.Role, error) {
	return roleService.repository.GetRole()
}

func (roleService *roleService) CreateRole(request model.CreateRoleRequest) (entity.Role, error) {
	role_request := entity.Role{
		Name: request.Name,
	}

	_, error := roleService.repository.CreateRole(role_request)

	return role_request, error
}
