package repository

import "svc-monitoring-maintenance/entity"

type HwReplacementStatusRepositoryInterface interface {
	GetHwReplacementStatus() ([]entity.HwReplacementStatus, error)
	CreateHwReplacementStatus(request *entity.HwReplacementStatus) ([]entity.HwReplacementStatus, error)
	UpdateHwReplacementStatus(request *entity.HwReplacementStatus) ([]entity.HwReplacementStatus, error)
	DeleteHwReplacementStatus(request *entity.HwReplacementStatus) error
}

func (repo *repository) GetHwReplacementStatus() ([]entity.HwReplacementStatus, error) {
	var status []entity.HwReplacementStatus

	error := repo.db.Raw("SELECT * FROM hw_replacement_status WHERE is_active = ? ORDER BY status ASC", "true").Find(&status).Error

	return status, error
}

func (repo *repository) CreateHwReplacementStatus(request *entity.HwReplacementStatus) ([]entity.HwReplacementStatus, error) {
	var status []entity.HwReplacementStatus

	error := repo.db.Table("hw_replacement_status").Create(&request).Find(&status).Error

	return status, error
}

func (repo *repository) UpdateHwReplacementStatus(request *entity.HwReplacementStatus) ([]entity.HwReplacementStatus, error) {
	var status []entity.HwReplacementStatus

	error := repo.db.Table("hw_replacement_status").Model(&request).Updates(request).Find(&status).Error

	return status, error
}

func (repo *repository) DeleteHwReplacementStatus(request *entity.HwReplacementStatus) error {
	var status []entity.HwReplacementStatus

	error := repo.db.Table("hw_replacement_status").Model(&request).Update("is_active", "false").Find(&status).Error

	return error
}
