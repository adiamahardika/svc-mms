package service

import (
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/model"
	"svc-monitoring-maintenance/repository"
)

type RoleServiceInterface interface {
	GetRole(request model.GetRoleRequest) ([]entity.Role, error)
	CreateRole(request model.CreateRoleRequest) (entity.Role, error)
	UpdateRole(request entity.Role) (entity.Role, error)
	DeleteRole(Id int) error
}

type roleService struct {
	repository repository.RoleRepositoryInteface
}

func RoleService(repository repository.RoleRepositoryInteface) *roleService {
	return &roleService{repository}
}

func (roleService *roleService) GetRole(request model.GetRoleRequest) ([]entity.Role, error) {
	return roleService.repository.GetRole(request)
}

func (roleService *roleService) CreateRole(request model.CreateRoleRequest) (entity.Role, error) {
	role_request := entity.Role{
		Name: request.Name,
	}

	_, error := roleService.repository.CreateRole(role_request)

	return role_request, error
}

func (roleService *roleService) UpdateRole(request entity.Role) (entity.Role, error) {
	role_request := entity.Role{
		Id:   request.Id,
		Name: request.Name,
	}

	_, error := roleService.repository.UpdateRole(role_request)

	return role_request, error
}

func (roleService *roleService) DeleteRole(Id int) error {

	error := roleService.repository.DeleteRole(Id)

	return error
}
