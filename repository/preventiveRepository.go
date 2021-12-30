package repository

import (
	"svc-monitoring-maintenance/entity"
)

type PreventiveRepositoryInterface interface {
	CreatePreventive(request entity.Preventive) (entity.Preventive, error)
}

func (repo *repository) CreatePreventive(request entity.Preventive) (entity.Preventive, error) {
	var preventive entity.Preventive

	error := repo.db.Table("preventive").Create(&request).Error

	return preventive, error
}
