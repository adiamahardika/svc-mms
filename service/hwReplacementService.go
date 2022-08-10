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

type HwReplacementServiceInterface interface {
	CreateHwReplacement(request *model.HwReplacementRequest, context *gin.Context) (entity.HwReplacement, error)
}

type hwReplacementService struct {
	repository repository.HwReplacementRepositoryInterface
}

func HwReplacementService(repository repository.HwReplacementRepositoryInterface) *hwReplacementService {
	return &hwReplacementService{repository}
}

func (hwReplacementService *hwReplacementService) CreateHwReplacement(request *model.HwReplacementRequest, context *gin.Context) (entity.HwReplacement, error) {
	var hw_replacement entity.HwReplacement
	date_now := time.Now()
	dir := os.Getenv("FILE_DIR")
	path := dir + "/hw_replacement/" + request.TicketCode + "/" + date_now.Format("2006-01-02")
	error := fmt.Errorf("error")
	attachment := ""

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
		new_request := &entity.HwReplacement{
			TicketCode:  request.TicketCode,
			HwId:        request.HwId,
			OldSN:       request.OldSN,
			NewSN:       request.NewSN,
			Description: request.Description,
			Attachment:  attachment,
			CreatedBy:   request.CreatedBy,
			CreatedAt:   date_now,
		}

		hw_replacement, error = hwReplacementService.repository.CreateHwReplacement(new_request)
	}

	return hw_replacement, error
}
