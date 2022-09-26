package repository

import "svc-monitoring-maintenance/entity"

type UserChecklistPreventiveRepositoryInterface interface {
	CreateUserChecklistPreventive(request []*entity.UserChecklistPreventive) ([]entity.UserChecklistPreventive, error)
}

func (repo *repository) CreateUserChecklistPreventive(request []*entity.UserChecklistPreventive) ([]entity.UserChecklistPreventive, error) {
	var userChecklistPreventive []entity.UserChecklistPreventive

	error := repo.db.Table("user_checklist_preventive").Create(&request).Find(&userChecklistPreventive).Error

	return userChecklistPreventive, error
}
