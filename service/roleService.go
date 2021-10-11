package service

import (
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/repository"
)

type RoleServiceInterface interface {
	GetRole() ([]entity.Role, error)
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