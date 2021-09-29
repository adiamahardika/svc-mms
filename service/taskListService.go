package service

import (
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/model"
	"svc-monitoring-maintenance/repository"
	"time"

	"github.com/gin-gonic/gin"
)

type TaskListServiceInterface interface {
	GetTaskList(request *model.GetTaskListRequest) ([]entity.TaskList, error)
	UpdateTaskList(request model.UpdateTaskListRequest, context *gin.Context)	(entity.TaskList, error)
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

func (taskListService *taskListService) UpdateTaskList(request model.UpdateTaskListRequest, context *gin.Context) (entity.TaskList, error) {
	date_now := time.Now()

	path := "D:/monitoring_maintenance/" + request.Attachment.Filename
	
	error_upload := context.SaveUploadedFile(request.Attachment, path)

	new_request := entity.TaskList {
		KodeTicket: request.KodeTicket,
		Description: request.Description,
		Attachment: path,
		TaskName: request.TaskName,
		Longitude: request.Latitude,
		Latitude: request.Latitude,
		AssignedBy: request.AssignedBy,
		Status: request.Status,
		CreatedAt: date_now,
	}

	ticket, error := taskListService.repository.UpdateTaskList(new_request)

	if (error_upload != nil) {
		
		return ticket, error_upload
		
	} else {

		return ticket, error
	}
}