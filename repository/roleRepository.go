package repository

import (
	"svc-monitoring-maintenance/entity"
)

type RoleRepositoryInteface interface {
	GetRole() ([]entity.Role, error)
	CreateRole(entity.Role) (entity.Role, error)
}

func (repo *repository) GetRole() ([]entity.Role, error) {
	var role []entity.Role

	error := repo.db.Raw("SELECT * FROM role ORDER BY name").Find(&role).Error

	return role, error
}

func (repo *repository) CreateRole(request entity.Role) (entity.Role, error) {
	var role entity.Role

	error := repo.db.Table("role").Create(&request).Error

	return role, error
}
