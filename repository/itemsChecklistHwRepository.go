package repository

import "svc-monitoring-maintenance/entity"

type ItemsChecklistHwRepositoryInterface interface {
	CreateItemsChecklistHw(request *entity.ItemsChecklistHw) ([]entity.ItemsChecklistHw, error)
	GetItemsChecklistHw() ([]entity.ItemsChecklistHw, error)
	UpdateItemsChecklistHw(request *entity.ItemsChecklistHw) (entity.ItemsChecklistHw, error)
}

func (repo *repository) CreateItemsChecklistHw(request *entity.ItemsChecklistHw) ([]entity.ItemsChecklistHw, error) {
	var itemsChecklistHw []entity.ItemsChecklistHw

	error := repo.db.Table("items_checklist_hw").Create(&request).Find(&itemsChecklistHw).Error

	return itemsChecklistHw, error
}

func (repo *repository) GetItemsChecklistHw() ([]entity.ItemsChecklistHw, error) {
	var itemsChecklistHw []entity.ItemsChecklistHw

	error := repo.db.Table("items_checklist_hw").Order("name ASC").Find(&itemsChecklistHw).Error

	return itemsChecklistHw, error
}

func (repo *repository) UpdateItemsChecklistHw(request *entity.ItemsChecklistHw) (entity.ItemsChecklistHw, error) {
	var itemsChecklistHw entity.ItemsChecklistHw

	error := repo.db.Table("items_checklist_hw").Model(&request).Updates(request).Find(&itemsChecklistHw).Error

	return itemsChecklistHw, error
}
