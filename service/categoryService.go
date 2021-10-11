package service

import (
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/repository"
)

type CategoryServiceInterface interface {
	GetCategory() ([]entity.Category, error)
}

type categoryService struct {
	repository repository.CategoryRepositoryInterface
}

func CategoryService(repository repository.CategoryRepositoryInterface) *categoryService {
	return &categoryService{repository}
}

func (categoryService *categoryService) GetCategory() ([]entity.Category, error) {
	return categoryService.repository.GetCategory()
}