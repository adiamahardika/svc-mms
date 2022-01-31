package repository

import (
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/model"
)

type AuthRepositoryInterface interface {
	Register(request model.RegisterRequest) (entity.User, error)
}

func (repo *repository) Register(request model.RegisterRequest) (entity.User, error) {
	var user entity.User

	error := repo.db.Table("users").Create(&request).Error

	return user, error
}
