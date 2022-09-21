package repository

import "svc-monitoring-maintenance/entity"

type MsChecklistHwRepositoryInterface interface {
	CreateMsChecklistHw(request *entity.MsChecklistHw) ([]entity.MsChecklistHw, error)
}

func (repo *repository) CreateMsChecklistHw(request *entity.MsChecklistHw) ([]entity.MsChecklistHw, error) {
	var msChecklistHw []entity.MsChecklistHw

	error := repo.db.Table("ms_checklist_hw").Create(&request).Find(&msChecklistHw).Error

	return msChecklistHw, error
}
