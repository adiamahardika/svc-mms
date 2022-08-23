package repository

import (
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/model"
)

type RoleRepositoryInteface interface {
	GetRole(request model.GetRoleRequest) ([]entity.Role, error)
	CreateRole(entity.Role) (entity.Role, error)
	UpdateRole(request entity.Role) (entity.Role, error)
	DeleteRole(Id int) error
}

func (repo *repository) GetRole(request model.GetRoleRequest) ([]entity.Role, error) {
	var role []entity.Role

	error := repo.db.Raw("SELECT role.*, JSON_AGG(JSON_BUILD_OBJECT('id', mms_web_permission.id, 'name', mms_web_permission.name, 'code', mms_web_permission.permission_code)) AS web_permission, JSON_AGG(JSON_BUILD_OBJECT('id', mms_app_permission.id, 'name', mms_app_permission.name, 'code', mms_app_permission.permission_code)) AS app_permission FROM role INNER JOIN  mms_role_has_web_permission ON (role.id = mms_role_has_web_permission.id_role) INNER JOIN mms_web_permission ON (mms_role_has_web_permission.id_permission = mms_web_permission.id) INNER JOIN  mms_role_has_app_permission ON (role.id = mms_role_has_app_permission.id_role) INNER JOIN mms_app_permission ON (mms_role_has_app_permission.id_permission = mms_app_permission.id)  WHERE is_active LIKE @IsActive GROUP BY role.id, mms_role_has_web_permission.id_role, mms_role_has_app_permission.id_role ORDER BY name", model.GetRoleRequest{
		IsActive: "%" + request.IsActive + "%",
	}).Find(&role).Error

	return role, error
}

func (repo *repository) CreateRole(request entity.Role) (entity.Role, error) {
	var role entity.Role

	error := repo.db.Table("role").Create(&request).Error

	return role, error
}

func (repo *repository) UpdateRole(request entity.Role) (entity.Role, error) {
	var role entity.Role

	error := repo.db.Raw("UPDATE role SET name = @Name WHERE id = @Id RETURNING role.*", request).Find(&role).Error

	return role, error
}

func (repo *repository) DeleteRole(Id int) error {
	var role entity.Role

	error := repo.db.Raw("UPDATE role SET is_active = ? WHERE id = ? RETURNING role.*", "false", Id).Find(&role).Error

	return error
}
