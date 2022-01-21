package repository

import "svc-monitoring-maintenance/entity"

type LogRepositoryInterface interface {
	CreateLog(request entity.LgServiceActivities) error
}

func (repo *repository) CreateLog(request entity.LgServiceActivities) error {
	error := repo.db.Table("lg_service_activities").Create(&request).Error

	return error
}
