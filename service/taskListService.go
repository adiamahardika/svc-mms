package service

import (
	"svc-ticket-monitoring/entity"
	"svc-ticket-monitoring/model"
	"svc-ticket-monitoring/repository"
)

type TaskListServiceInterface interface {
	FindTaskList(request *model.GetTaskListRequest) ([]entity.TaskList, error)
}

type taskListService struct {
	repository repository.TaskListRepositoryInterface
}

func TaskListService(repository repository.TaskListRepositoryInterface) *taskListService {
	return &taskListService{repository}
}

func (taskListService *taskListService) FindTaskList(request *model.GetTaskListRequest) ([]entity.TaskList, error) {
	return taskListService.repository.FindTaskList(request)
}