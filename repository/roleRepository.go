package repository

import (
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/model"
)

type RoleRepositoryInteface interface {
	GetRole(request *model.GetRoleRequest) ([]entity.Role, error)
	CreateRole(request *entity.Role) ([]entity.Role, error)
	UpdateRole(request *entity.Role) (entity.Role, error)
	DeleteRole(id *int) error
}

func (repo *repository) GetRole(request *model.GetRoleRequest) ([]entity.Role, error) {
	var role []entity.Role

	error := repo.db.Raw("SELECT role.*, JSON_AGG(DISTINCT mms_web_permission.*) AS web_permission, JSON_AGG(DISTINCT mms_app_permission.*) AS app_permission FROM role INNER JOIN  mms_role_has_web_permission ON (role.id = mms_role_has_web_permission.id_role) INNER JOIN mms_web_permission ON (mms_role_has_web_permission.id_permission = mms_web_permission.id) INNER JOIN  mms_role_has_app_permission ON (role.id = mms_role_has_app_permission.id_role) INNER JOIN mms_app_permission ON (mms_role_has_app_permission.id_permission = mms_app_permission.id)  WHERE is_active LIKE @IsActive GROUP BY mms_role_has_web_permission.id_role, mms_role_has_app_permission.id_role, role.id ORDER BY name", model.GetRoleRequest{
		IsActive: "%" + request.IsActive + "%",
	}).Find(&role).Error

	return role, error
}

func (repo *repository) CreateRole(request *entity.Role) ([]entity.Role, error) {
	var role []entity.Role

	error := repo.db.Raw("INSERT INTO role(name, is_active) VALUES(@Name, @IsActive) RETURNING role.*", request).Find(&role).Error

	return role, error
}

func (repo *repository) UpdateRole(request *entity.Role) (entity.Role, error) {
	var role entity.Role

	error := repo.db.Raw("UPDATE role SET name = @Name WHERE id = @Id RETURNING role.*", request).Find(&role).Error

	return role, error
}

func (repo *repository) DeleteRole(id *int) error {
	var role entity.Role

	error := repo.db.Raw("UPDATE role SET is_active = ? WHERE id = ? RETURNING role.*", "false", id).Find(&role).Error

	return error
}
