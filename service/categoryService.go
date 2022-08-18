package service

import (
	"encoding/json"
	"math"
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/model"
	"svc-monitoring-maintenance/repository"
	"time"
)

type CategoryServiceInterface interface {
	GetCategory(request *model.GetCategoryRequest) ([]model.CreateCategoryRequest, float64, error)
	CreateCategory(request *model.CreateCategoryRequest) (model.CreateCategoryRequest, error)
	UpdateCategory(request *model.CreateCategoryRequest) (model.CreateCategoryRequest, error)
	DeleteCategory(id *int) error
	GetDetailCategory(request *string) ([]model.CreateCategoryRequest, error)
}

type categoryService struct {
	categoryRepository    repository.CategoryRepositoryInterface
	subCategoryRepository repository.SubCategoryRepositoryInterface
}

func CategoryService(categoryRepository repository.CategoryRepositoryInterface, subCategoryRepository repository.SubCategoryRepositoryInterface) *categoryService {
	return &categoryService{categoryRepository, subCategoryRepository}
}

func (categoryService *categoryService) GetCategory(request *model.GetCategoryRequest) ([]model.CreateCategoryRequest, float64, error) {
	var response []model.CreateCategoryRequest

	if request.Size == 0 {
		request.Size = math.MaxInt16
	}
	request.StartIndex = request.PageNo * request.Size
	total_data, error := categoryService.categoryRepository.CountCategory(request)

	total_pages := math.Ceil(float64(total_data) / float64(request.Size))

	category, error := categoryService.categoryRepository.GetCategory(request)

	for _, value := range category {
		var sub_category []entity.SubCategory
		json.Unmarshal([]byte(value.SubCategory), &sub_category)

		response = append(response, model.CreateCategoryRequest{
			Id:          value.Id,
			Name:        value.Name,
			SubCategory: sub_category,
			IsActive:    value.IsActive,
			UpdateAt:    value.UpdateAt,
		})
	}

	return response, total_pages, error
}

func (categoryService *categoryService) CreateCategory(request *model.CreateCategoryRequest) (model.CreateCategoryRequest, error) {
	var request_sc []*entity.SubCategory
	var response_sc []entity.SubCategory
	var response model.CreateCategoryRequest
	date_now := time.Now()

	request.UpdateAt = date_now
	request.IsActive = "true"

	category, error := categoryService.categoryRepository.CreateCategory(request)

	if error == nil {
		for _, value := range request.SubCategory {
			request_sc = append(request_sc, &entity.SubCategory{
				Name:       value.Name,
				IdCategory: category.Id,
				Priority:   value.Priority,
				CreatedAt:  date_now,
				UpdatedAt:  date_now,
			})
		}
		response_sc, error = categoryService.subCategoryRepository.CreateSubCategory(request_sc)

		response = model.CreateCategoryRequest{
			Id:          category.Id,
			Name:        category.Name,
			SubCategory: response_sc,
			IsActive:    category.IsActive,
			UpdateAt:    category.UpdateAt,
		}
	}

	return response, error
}

func (categoryService *categoryService) UpdateCategory(request *model.CreateCategoryRequest) (model.CreateCategoryRequest, error) {
	var request_sc []*entity.SubCategory
	var response_sc []entity.SubCategory
	date_now := time.Now()

	request.UpdateAt = date_now
	category, error := categoryService.categoryRepository.UpdateCategory(request)

	if error == nil {
		error = categoryService.subCategoryRepository.DeleteSubCategory(&request.Id)

		if error == nil {
			for _, value := range request.SubCategory {
				request_sc = append(request_sc, &entity.SubCategory{
					Name:       value.Name,
					IdCategory: request.Id,
					Priority:   value.Priority,
					CreatedAt:  date_now,
					UpdatedAt:  date_now,
				})
			}

			response_sc, error = categoryService.subCategoryRepository.CreateSubCategory(request_sc)
			category.SubCategory = response_sc
		}
	}

	return category, error
}

func (categoryService *categoryService) DeleteCategory(id *int) error {

	error := categoryService.categoryRepository.DeleteCategory(id)

	return error
}

func (categoryService *categoryService) GetDetailCategory(request *string) ([]model.CreateCategoryRequest, error) {
	var response []model.CreateCategoryRequest

	category, error := categoryService.categoryRepository.GetDetailCategory(request)

	for _, value := range category {
		var sub_category []entity.SubCategory
		json.Unmarshal([]byte(value.SubCategory), &sub_category)

		response = append(response, model.CreateCategoryRequest{
			Id:          value.Id,
			Name:        value.Name,
			SubCategory: sub_category,
			IsActive:    value.IsActive,
			UpdateAt:    value.UpdateAt,
		})
	}

	return response, error
}
