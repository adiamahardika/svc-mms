package service

import (
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/model"
	"svc-monitoring-maintenance/repository"
)

type TaskListServiceInterface interface {
	GetTaskList(request *model.GetTaskListRequest) ([]entity.TaskList, error)
}

type taskListService struct {
	repository repository.TaskListRepositoryInterface
}

func TaskListService(repository repository.TaskListRepositoryInterface) *taskListService {
	return &taskListService{repository}
}

func (taskListService *taskListService) GetTaskList(request *model.GetTaskListRequest) ([]entity.TaskList, error) {
	return taskListService.repository.GetTaskList(request)
}