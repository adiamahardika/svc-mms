package repository

import "svc-monitoring-maintenance/entity"

type WebPermissionRepositoryInterface interface {
	GetWebPermission() ([]entity.MmsWebPermission, error)
}

func (repo *repository) GetWebPermission() ([]entity.MmsWebPermission, error) {
	var web_permission []entity.MmsWebPermission

	error := repo.db.Raw("SELECT * FROM mms_web_permission ORDER BY name ASC").Find(&web_permission).Error

	return web_permission, error
}
