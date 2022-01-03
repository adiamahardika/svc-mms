package repository

import "svc-monitoring-maintenance/entity"

type TaskPreventiveRepositoryInterface interface {
	UpdateTaskPreventive(request entity.TaskPreventive) (entity.TaskPreventive, error)
}

func (repo *repository) UpdateTaskPreventive(request entity.TaskPreventive) (entity.TaskPreventive, error) {
	var task_preventive entity.TaskPreventive

	error := repo.db.Table("task_preventive").Create(&request).Error

	return task_preventive, error
}
