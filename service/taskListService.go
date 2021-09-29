package service

import (
	"os"
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
	
	task_list, error := taskListService.repository.GetTaskList(request)

	host := os.Getenv("DB_HOST")
	port := os.Getenv("PORT")

	for index := range task_list {
		task_list[index].Attachment = "http://" + host + port + "/assets/" + task_list[index].Attachment
	}

	return task_list, error
}

func (taskListService *taskListService) UpdateTaskList(request model.UpdateTaskListRequest, context *gin.Context) (entity.TaskList, error) {
	date_now := time.Now()
	
	error_upload := context.SaveUploadedFile(request.Attachment, request.Attachment.Filename)

	new_request := entity.TaskList {
		KodeTicket: request.KodeTicket,
		Description: request.Description,
		Attachment: request.Attachment.Filename,
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