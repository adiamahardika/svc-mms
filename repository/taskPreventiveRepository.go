package repository

import (
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/model"
)

type TaskPreventiveRepositoryInterface interface {
	UpdateTaskPreventive(request entity.TaskPreventive) (entity.TaskPreventive, error)
	GetTaskPreventive(request *model.GetTaskPreventiveRequest) ([]model.GetTaskPreventiveResponse, error)
}

func (repo *repository) UpdateTaskPreventive(request entity.TaskPreventive) (entity.TaskPreventive, error) {
	var task_preventive entity.TaskPreventive

	error := repo.db.Table("task_preventive").Create(&request).Error

	return task_preventive, error
}

func (repo *repository) GetTaskPreventive(request *model.GetTaskPreventiveRequest) ([]model.GetTaskPreventiveResponse, error) {
	var task_preventive []model.GetTaskPreventiveResponse

	error := repo.db.Raw("SELECT task_preventive.*, users.name as user_name FROM task_preventive LEFT OUTER JOIN users ON (task_preventive.assigned_by = CAST(users.id AS varchar(10))) WHERE prev_code LIKE @PrevCode ORDER BY index", model.GetTaskPreventiveRequest{
		PrevCode: "%" + request.PrevCode + "%",
	}).Find(&task_preventive).Error

	return task_preventive, error
}
