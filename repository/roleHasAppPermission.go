package repository

import (
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/model"
)

type RoleHasAppPermissionRepositoryInterface interface {
	CreateRoleHasAppPermission(request []*model.CreateRoleHasAppPermissionRequest) error
	DeleteRoleHasAppPermission(id_role *int) error
}

func (repo *repository) CreateRoleHasAppPermission(request []*model.CreateRoleHasAppPermissionRequest) error {

	error := repo.db.Table("mms_role_has_app_permission").Create(&request).Error

	return error
}

func (repo *repository) DeleteRoleHasAppPermission(id_role *int) error {

	var role_has_permission *entity.RoleHasAppPermission

	error := repo.db.Raw("DELETE FROM mms_role_has_app_permission WHERE id_role = ? RETURNING mms_role_has_app_permission.*", id_role).Find(&role_has_permission).Error

	return error
}
