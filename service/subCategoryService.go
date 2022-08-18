package service

import (
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/repository"
)

type SubCategoryServiceInterface interface {
	GetSubCategory() ([]entity.SubCategory, error)
}

type subCategoryService struct {
	subCategoryRepository repository.SubCategoryRepositoryInterface
}

func SubCategoryService(subCategoryRepository repository.SubCategoryRepositoryInterface) *subCategoryService {
	return &subCategoryService{subCategoryRepository}
}

func (subCategoryService *subCategoryService) GetSubCategory() ([]entity.SubCategory, error) {

	sub_category, error := subCategoryService.subCategoryRepository.GetSubCategory()

	return sub_category, error
}
