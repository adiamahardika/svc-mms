package repository

import (
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/model"
)

type HwReplacementRepositoryInterface interface {
	CreateHwReplacement(request *entity.HwReplacement) (entity.HwReplacement, error)
	GetHwReplacement(request *model.GetHwReplacementRequest) ([]entity.HwReplacement, error)
}

func (repo *repository) CreateHwReplacement(request *entity.HwReplacement) (entity.HwReplacement, error) {
	var hw_replacement entity.HwReplacement

	error := repo.db.Table("hw_replacement").Create(&request).Error

	return hw_replacement, error
}

func (repo *repository) GetHwReplacement(request *model.GetHwReplacementRequest) ([]entity.HwReplacement, error) {
	var hw_replacement []entity.HwReplacement

	error := repo.db.Raw("SELECT * FROM hw_replacement WHERE ticket_code = @TicketCode ORDER BY  created_at ASC", request).Find(&hw_replacement).Error

	return hw_replacement, error
}
