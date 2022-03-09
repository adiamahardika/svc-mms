package repository

import (
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/model"
)

type TaskListRepositoryInterface interface {
	GetTaskList(request *model.GetTaskListRequest) ([]model.GetTaskListResponse, error)
	UpdateTaskList(request entity.TaskList) (entity.TaskList, error)
}

func (repo *repository) GetTaskList(request *model.GetTaskListRequest) ([]model.GetTaskListResponse, error) {
	var task_list []model.GetTaskListResponse

	error := repo.db.Raw("SELECT task_list.*, users.name as user_name FROM task_list LEFT OUTER JOIN users ON (task_list.assigned_by = CAST(users.id AS varchar(10))) WHERE ticket_code = @TicketCode ORDER BY index", model.GetTaskListRequest{
		TicketCode: request.TicketCode,
	}).Find(&task_list).Error

	return task_list, error
}

func (repo *repository) UpdateTaskList(request entity.TaskList) (entity.TaskList, error) {
	var task_list entity.TaskList

	error := repo.db.Table("task_list").Create(&request).Error

	return task_list, error
}
