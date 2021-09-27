package repository

import "svc-monitoring-maintenance/entity"

type RoleRepository interface {
	GetRole() ([]entity.Role, error)
}

func (repo *repository) GetRole() ([]entity.Role, error) {
	var role []entity.Role

	error := repo.db.Raw("SELECT * FROM role ORDER BY name").Find(&role).Error

	return role, error
}