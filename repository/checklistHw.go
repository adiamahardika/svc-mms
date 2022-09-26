package repository

import (
	"svc-monitoring-maintenance/entity"
)

type ChecklistHwRepositoryInterface interface {
	CreateChecklistHw(request []*entity.ChecklistHw) ([]entity.ChecklistHw, error)
	GetChecklistHw(request *string) ([]entity.ChecklistHw, error)
}

func (repo *repository) CreateChecklistHw(request []*entity.ChecklistHw) ([]entity.ChecklistHw, error) {
	var checklistHw []entity.ChecklistHw

	error := repo.db.Table("checklist_hw").Create(&request).Find(&checklistHw).Error

	return checklistHw, error
}

func (repo *repository) GetChecklistHw(request *string) ([]entity.ChecklistHw, error) {
	var checklistHw []entity.ChecklistHw

	error := repo.db.Table("checklist_hw").Where("prev_code = ?", request).Find(&checklistHw).Error

	return checklistHw, error
}
