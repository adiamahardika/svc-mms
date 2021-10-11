package repository

import "svc-monitoring-maintenance/entity"

type CategoryRepositoryInterface interface {
	GetCategory() ([]entity.Category, error)
}

func (repo *repository) GetCategory() ([]entity.Category, error) {
	var category []entity.Category

	error := repo.db.Raw("SELECT * FROM category ORDER BY name").Find(&category).Error

	return category, error
}