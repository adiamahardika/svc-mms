package repository

import (
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/model"
)

type RoleHasWebPermissionRepositoryInterface interface {
	CreateRoleHasWebPermission(request []*model.CreateRoleHasWebPermissionRequest) error
	DeleteRoleHasWebPermission(id_role *int) error
}

func (repo *repository) CreateRoleHasWebPermission(request []*model.CreateRoleHasWebPermissionRequest) error {

	error := repo.db.Table("mms_role_has_web_permission").Create(&request).Error

	return error
}

func (repo *repository) DeleteRoleHasWebPermission(id_role *int) error {

	var role_has_permission *entity.RoleHasWebPermission

	error := repo.db.Raw("DELETE FROM mms_role_has_web_permission WHERE id_role = ? RETURNING mms_role_has_web_permission.*", id_role).Find(&role_has_permission).Error

	return error
}
