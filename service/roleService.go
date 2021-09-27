package service

import (
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/repository"
)

type RoleServiceInterface interface {
	GetRole() ([]entity.Role, error)
}

type roleService struct {
	repository repository.RoleRepository
}

func RoleService(repository repository.RoleRepository) *roleService {
	return &roleService{repository}
}

func (roleService *roleService) GetRole() ([]entity.Role, error) {
	return roleService.repository.GetRole()
}