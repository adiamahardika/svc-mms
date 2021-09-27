package repository

import (
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/model"
)

type TaskListRepositoryInterface interface {
	GetTaskList(request *model.GetTaskListRequest) ([]entity.TaskList, error)
}

func (repo *repository) GetTaskList(request *model.GetTaskListRequest) ([]entity.TaskList, error) {
	var task_list []entity.TaskList

	error := repo.db.Raw("SELECT * FROM task_list WHERE kode_ticket LIKE @KodeTicket", model.GetTaskListRequest{
		KodeTicket: "%" + request.KodeTicket + "%",
	}).Find(&task_list).Error

	return task_list, error
}