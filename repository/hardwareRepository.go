package repository

import (
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/model"
)

type HardwareRepositoryInterface interface {
	GetHardware(request *model.GetHardwareRequest) ([]entity.Hardware, error)
}

func (repo *repository) GetHardware(request *model.GetHardwareRequest) ([]entity.Hardware, error) {

	var hardware []entity.Hardware

	error := repo.db.Raw("SELECT * FROM hardware WHERE name LIKE @Search AND is_active LIKE @IsActive ORDER BY name ASC", model.GetHardwareRequest{
		Search:   "%" + request.Search + "%",
		IsActive: "%" + request.IsActive + "%",
	}).Find(&hardware).Error

	return hardware, error
}
