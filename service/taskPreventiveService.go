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

type TaskPreventiveServiceInterface interface {
	UpdateTaskPreventive(request model.UpdateTaskPreventiveRequest, context *gin.Context) (entity.TaskPreventive, error)
}

type taskPreventiveService struct {
	repository repository.TaskPreventiveRepositoryInterface
}

func TaskPreventiveService(repository repository.TaskPreventiveRepositoryInterface) *taskPreventiveService {
	return &taskPreventiveService{repository}
}

func (taskPreventiveService *taskPreventiveService) UpdateTaskPreventive(request model.UpdateTaskPreventiveRequest, context *gin.Context) (entity.TaskPreventive, error) {
	var task_preventive entity.TaskPreventive
	date_now := time.Now()
	dir := os.Getenv("FILE_DIR")
	path := dir + "/preventive/" + date_now.Format("2006-01-02") + "/" + request.PrevCode
	error := fmt.Errorf("error")
	attachment := "-"

	_, check_dir_error := os.Stat(path)

	if os.IsNotExist(check_dir_error) {
		check_dir_error := os.MkdirAll(path, 0755)

		if check_dir_error != nil {
			error = check_dir_error
		}
	}

	if request.Attachment != nil {
		attachment = request.Attachment.Filename
		error = context.SaveUploadedFile(request.Attachment, path+"/"+attachment)
	} else {
		error = nil
	}

	if error == nil {
		new_request := entity.TaskPreventive{
			PrevCode:    request.PrevCode,
			Description: request.Description,
			Attachment:  attachment,
			TaskName:    request.TaskName,
			Longitude:   request.Longitude,
			Latitude:    request.Latitude,
			AssignedBy:  request.AssignedBy,
			Status:      request.Status,
			Index:       request.Index,
			CreatedAt:   date_now,
		}

		task_preventive, error = taskPreventiveService.repository.UpdateTaskPreventive(new_request)
	}

	return task_preventive, error
}
