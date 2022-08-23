package service

import (
	"encoding/json"
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/model"
	"svc-monitoring-maintenance/repository"
)

type RoleServiceInterface interface {
	GetRole(request model.GetRoleRequest) ([]model.GetRoleResponse, error)
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

func (roleService *roleService) GetRole(request model.GetRoleRequest) ([]model.GetRoleResponse, error) {
	var response []model.GetRoleResponse
	role, error := roleService.repository.GetRole(request)

	for _, value := range role {
		var web_permission []*entity.MmsWebPermission
		var app_permission []*entity.MmsAppPermission
		json.Unmarshal([]byte(value.WebPermission), &web_permission)
		json.Unmarshal([]byte(value.AppPermission), &app_permission)

		response = append(response, model.GetRoleResponse{
			Name:          value.Name,
			Id:            value.Id,
			WebPermission: web_permission,
			AppPermission: app_permission,
		})
	}

	return response, error
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
