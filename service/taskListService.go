package service

import (
	"fmt"
	"os"
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/model"
	"svc-monitoring-maintenance/repository"
	"time"

	"github.com/gin-gonic/gin"
)

type TaskListServiceInterface interface {
	GetTaskList(request *model.GetTaskListRequest) ([]model.GetTaskListResponse, error)
	UpdateTaskList(request model.UpdateTaskListRequest, context *gin.Context)	(entity.TaskList, error)
}

type taskListService struct {
	repository repository.TaskListRepositoryInterface
}

func TaskListService(repository repository.TaskListRepositoryInterface) *taskListService {
	return &taskListService{repository}
}

func (taskListService *taskListService) GetTaskList(request *model.GetTaskListRequest) ([]model.GetTaskListResponse, error) {
	
	task_list, error := taskListService.repository.GetTaskList(request)

	url := os.Getenv("FILE_URL")

	for index := range task_list {
		date := task_list[index].CreatedAt.Format("2006-01-02")
		ticket_code := task_list[index].TicketCode
		file_name := task_list[index].Attachment
		task_list[index].Attachment = url + date + "/" + ticket_code + "/" + file_name
	}

	return task_list, error
}

func (taskListService *taskListService) UpdateTaskList(request model.UpdateTaskListRequest, context *gin.Context) (entity.TaskList, error) {
	var ticket entity.TaskList
	date_now := time.Now()
	dir := os.Getenv("FILE_DIR")
	path := dir + date_now.Format("2006-01-02") + "/" + request.TicketCode 
	error := fmt.Errorf("error")

	_, check_dir_error := os.Stat(path)

	if (os.IsNotExist(check_dir_error)) {
		check_dir_error := os.MkdirAll(path, 0755)
		
		if (check_dir_error != nil) {
			error = check_dir_error
		}
	}

	error = context.SaveUploadedFile(request.Attachment, path + "/" + request.Attachment.Filename)
	if (error == nil) {
		new_request := entity.TaskList {
			TicketCode: request.TicketCode,
			Description: request.Description,
			Attachment: request.Attachment.Filename,
			TaskName: request.TaskName,
			Longitude: request.Latitude,
			Latitude: request.Latitude,
			AssignedBy: request.AssignedBy,
			Status: request.Status,
			Index: request.Index,
			CreatedAt: date_now,
		}
	
		ticket, error = taskListService.repository.UpdateTaskList(new_request)
	}

	return ticket, error
}