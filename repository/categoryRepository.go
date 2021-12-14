package repository

import (
	"svc-monitoring-maintenance/entity"
)

type CategoryRepositoryInterface interface {
	GetCategory() ([]entity.Category, error)
	CreateCategory(request entity.Category) (entity.Category, error)
}

func (repo *repository) GetCategory() ([]entity.Category, error) {
	var category []entity.Category

	error := repo.db.Raw("SELECT * FROM category ORDER BY name").Find(&category).Error

	return category, error
}

func (repo *repository) CreateCategory(request entity.Category) (entity.Category, error) {
	var category entity.Category

	error := repo.db.Table("category").Create(&request).Error

	return category, error
}