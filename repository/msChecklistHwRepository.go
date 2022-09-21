package repository

import "svc-monitoring-maintenance/entity"

type MsChecklistHwRepositoryInterface interface {
	CreateMsChecklistHw(request *entity.MsChecklistHw) ([]entity.MsChecklistHw, error)
	GetMsChecklistHw() ([]entity.MsChecklistHw, error)
}

func (repo *repository) CreateMsChecklistHw(request *entity.MsChecklistHw) ([]entity.MsChecklistHw, error) {
	var msChecklistHw []entity.MsChecklistHw

	error := repo.db.Table("ms_checklist_hw").Create(&request).Find(&msChecklistHw).Error

	return msChecklistHw, error
}

func (repo *repository) GetMsChecklistHw() ([]entity.MsChecklistHw, error) {
	var msChecklistHw []entity.MsChecklistHw

	error := repo.db.Table("ms_checklist_hw").Order("name ASC").Find(&msChecklistHw).Error

	return msChecklistHw, error
}
