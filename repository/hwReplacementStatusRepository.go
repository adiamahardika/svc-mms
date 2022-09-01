package repository

import "svc-monitoring-maintenance/entity"

type HwReplacementStatusRepositoryInterface interface {
	GetHwReplacementStatus() ([]entity.HwReplacementStatus, error)
}

func (repo *repository) GetHwReplacementStatus() ([]entity.HwReplacementStatus, error) {
	var status []entity.HwReplacementStatus

	error := repo.db.Raw("SELECT * FROM hw_replacement_status WHERE is_active = ? ORDER BY status ASC", "true").Find(&status).Error

	return status, error
}
