package service

import (
	"encoding/json"
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/model"
	"svc-monitoring-maintenance/repository"
)

type RoleServiceInterface interface {
	GetRole(request *model.GetRoleRequest) ([]model.GetRoleResponse, error)
	CreateRole(request *model.CreateRoleRequest) ([]entity.Role, error)
	UpdateRole(request *entity.Role) (entity.Role, error)
	DeleteRole(id *int) error
}

type roleService struct {
	roleRepository repository.RoleRepositoryInteface
	rhwpRepository repository.RoleHasWebPermissionRepositoryInterface
	rhapRepository repository.RoleHasAppPermissionRepositoryInterface
}

func RoleService(roleRepository repository.RoleRepositoryInteface, rhwpRepository repository.RoleHasWebPermissionRepositoryInterface,
	rhapRepository repository.RoleHasAppPermissionRepositoryInterface) *roleService {
	return &roleService{roleRepository, rhwpRepository, rhapRepository}
}

func (roleService *roleService) GetRole(request *model.GetRoleRequest) ([]model.GetRoleResponse, error) {
	var response []model.GetRoleResponse
	role, error := roleService.roleRepository.GetRole(request)

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

func (roleService *roleService) CreateRole(request *model.CreateRoleRequest) ([]entity.Role, error) {
	var rhwp_request []*model.CreateRoleHasWebPermissionRequest
	var rhap_request []*model.CreateRoleHasAppPermissionRequest

	role, error := roleService.roleRepository.CreateRole(&entity.Role{
		Name:     request.Name,
		IsActive: "true",
	})

	if error == nil {
		for _, value := range request.WebPermission {
			rhwp_request = append(rhwp_request, &model.CreateRoleHasWebPermissionRequest{IdRole: role[0].Id, IdPermission: value.Id})
		}
		error = roleService.rhwpRepository.CreateRoleHasWebPermission(rhwp_request)
	}
	if error == nil {
		for _, value := range request.AppPermission {
			rhap_request = append(rhap_request, &model.CreateRoleHasAppPermissionRequest{IdRole: role[0].Id, IdPermission: value.Id})
		}
		error = roleService.rhapRepository.CreateRoleHasAppPermission(rhap_request)
	}

	return role, error
}

func (roleService *roleService) UpdateRole(request *entity.Role) (entity.Role, error) {
	role_request := &entity.Role{
		Id:   request.Id,
		Name: request.Name,
	}

	_, error := roleService.roleRepository.UpdateRole(role_request)

	return *role_request, error
}

func (roleService *roleService) DeleteRole(id *int) error {

	error := roleService.roleRepository.DeleteRole(id)

	return error
}
