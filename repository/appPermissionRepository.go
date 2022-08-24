package repository

import "svc-monitoring-maintenance/entity"

type AppPermissionRepositoryInterface interface {
	GetAppPermission() ([]entity.MmsAppPermission, error)
}

func (repo *repository) GetAppPermission() ([]entity.MmsAppPermission, error) {
	var app_permission []entity.MmsAppPermission

	error := repo.db.Raw("SELECT * FROM mms_app_permission ORDER BY name ASC").Find(&app_permission).Error

	return app_permission, error
}
