package repository

import (
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/model"
)

type CategoryRepositoryInterface interface {
	GetCategory(request *model.GetCategoryRequest) ([]entity.Category, error)
	CountCategory(request *model.GetCategoryRequest) (int, error)
	CreateCategory(request *model.CreateCategoryRequest) (model.CreateCategoryRequest, error)
	UpdateCategory(request *model.CreateCategoryRequest) (model.CreateCategoryRequest, error)
	DeleteCategory(id *int) error
	GetDetailCategory(request *string) ([]entity.Category, error)
}

func (repo *repository) GetCategory(request *model.GetCategoryRequest) ([]entity.Category, error) {
	var category []entity.Category

	error := repo.db.Raw("SELECT mms_category.*, JSON_AGG(JSON_BUILD_OBJECT('id', mms_sub_category.id, 'idCategory', mms_sub_category.id_category, 'name', mms_sub_category.name, 'priority', mms_sub_category.priority)) AS sub_category FROM mms_category LEFT OUTER JOIN mms_sub_category ON (mms_category.id = mms_sub_category.id_category) WHERE is_active LIKE @IsActive GROUP BY mms_category.id ORDER BY name ASC LIMIT @Size OFFSET @StartIndex", model.GetCategoryRequest{
		IsActive:   "%" + request.IsActive + "%",
		Size:       request.Size,
		StartIndex: request.StartIndex,
	}).Find(&category).Error

	return category, error
}

func (repo *repository) CountCategory(request *model.GetCategoryRequest) (int, error) {
	var total_data int

	error := repo.db.Raw("SELECT COUNT(*) as total_data FROM mms_category WHERE is_active LIKE @IsActive", model.GetCategoryRequest{
		IsActive: "%" + request.IsActive + "%",
	}).Find(&total_data).Error

	return total_data, error
}

func (repo *repository) CreateCategory(request *model.CreateCategoryRequest) (model.CreateCategoryRequest, error) {
	var category model.CreateCategoryRequest

	error := repo.db.Raw("INSERT INTO mms_category(name, is_active, update_at) VALUES(@Name, @IsActive, @UpdateAt) RETURNING mms_category.*", request).Find(&category).Error

	return category, error
}

func (repo *repository) UpdateCategory(request *model.CreateCategoryRequest) (model.CreateCategoryRequest, error) {

	var category model.CreateCategoryRequest

	error := repo.db.Raw("UPDATE mms_category SET name = @Name, update_at = @UpdateAt WHERE id = @Id RETURNING mms_category.*", request).Find(&category).Error

	return category, error
}

func (repo *repository) DeleteCategory(id *int) error {
	var category entity.Category

	error := repo.db.Raw("UPDATE mms_category SET is_active = ? WHERE id = ? RETURNING mms_category.*", "false", id).Find(&category).Error

	return error
}

func (repo *repository) GetDetailCategory(request *string) ([]entity.Category, error) {
	var category []entity.Category

	error := repo.db.Raw("SELECT mms_category.*, JSON_AGG(JSON_BUILD_OBJECT('id', mms_sub_category.id, 'idCategory', mms_sub_category.id_category, 'name', mms_sub_category.name, 'priority', mms_sub_category.priority)) AS sub_category FROM mms_category LEFT OUTER JOIN mms_sub_category ON (mms_category.id = mms_sub_category.id_category) WHERE mms_category.id = ? GROUP BY mms_category.id", request).Find(&category).Error

	return category, error
}
