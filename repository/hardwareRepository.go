package repository

import (
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/model"
)

type HardwareRepositoryInterface interface {
	GetHardware(request *model.GetHardwareRequest) ([]entity.Hardware, error)
	CreateHardware(request *entity.Hardware) (entity.Hardware, error)
	UpdateHardware(request *entity.Hardware) (entity.Hardware, error)
	DeleteHardware(id *int) error
}

func (repo *repository) GetHardware(request *model.GetHardwareRequest) ([]entity.Hardware, error) {

	var hardware []entity.Hardware

	error := repo.db.Raw("SELECT * FROM hardware WHERE name LIKE @Search AND is_active LIKE @IsActive ORDER BY name ASC", model.GetHardwareRequest{
		Search:   "%" + request.Search + "%",
		IsActive: "%" + request.IsActive + "%",
	}).Find(&hardware).Error

	return hardware, error
}

func (repo *repository) CreateHardware(request *entity.Hardware) (entity.Hardware, error) {

	var hardware entity.Hardware

	error := repo.db.Table("hardware").Create(&request).Error

	return hardware, error
}

func (repo *repository) UpdateHardware(request *entity.Hardware) (entity.Hardware, error) {
	var hardware entity.Hardware

	error := repo.db.Raw("UPDATE hardware SET name = @Name WHERE id = @Id RETURNING hardware.*", request).Find(&hardware).Error

	return hardware, error
}

func (repo *repository) DeleteHardware(id *int) error {
	var hardware entity.Hardware

	error := repo.db.Raw("UPDATE hardware SET is_active = ? WHERE id = ? RETURNING hardware.*", "false", id).Find(&hardware).Error

	return error
}
