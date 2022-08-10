package repository

import (
	"svc-monitoring-maintenance/entity"
)

type HwReplacementRepositoryInterface interface {
	CreateHwReplacement(request *entity.HwReplacement) (entity.HwReplacement, error)
}

func (repo *repository) CreateHwReplacement(request *entity.HwReplacement) (entity.HwReplacement, error) {
	var hw_replacement entity.HwReplacement

	error := repo.db.Table("hw_replacement").Create(&request).Error

	return hw_replacement, error
}
