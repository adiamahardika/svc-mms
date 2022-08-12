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
	CreateHwReplacement(request *model.CreateHwReplacementRequest, context *gin.Context) (model.GetHwReplacementResponse, error)
	GetHwReplacement(request *model.GetHwReplacementRequest) (model.GetHwReplacementResponse, error)
}

type hwReplacementService struct {
	hwReplacementRepository repository.HwReplacementRepositoryInterface
	ticketRepository        repository.TicketRepositoryInterface
}

func HwReplacementService(hwReplacementRepository repository.HwReplacementRepositoryInterface, ticketRepository repository.TicketRepositoryInterface) *hwReplacementService {
	return &hwReplacementService{hwReplacementRepository, ticketRepository}
}

func (hwReplacementService *hwReplacementService) CreateHwReplacement(request *model.CreateHwReplacementRequest, context *gin.Context) (model.GetHwReplacementResponse, error) {
	var response model.GetHwReplacementResponse
	var hw_replacement []entity.HwReplacement

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

		_, error = hwReplacementService.hwReplacementRepository.CreateHwReplacement(new_request)
	}

	if error == nil {

		ticket, error := hwReplacementService.ticketRepository.CheckTicketCode(request.TicketCode)

		if error == nil && len(ticket) > 0 {
			response.TicketCode = request.TicketCode
			response.NoSPM = ticket[0].NoSPM
			response.NoReqSPM = ticket[0].NoReqSPM

			hw_replacement, error = hwReplacementService.hwReplacementRepository.GetHwReplacement(&model.GetHwReplacementRequest{
				TicketCode: request.TicketCode,
			})

			url := os.Getenv("FILE_URL")

			for index := range hw_replacement {
				date := hw_replacement[index].CreatedAt.Format("2006-01-02")
				file_name := hw_replacement[index].Attachment
				hw_replacement[index].Attachment = url + "hw_replacement/" + request.TicketCode + "/" + date + "/" + file_name
			}

			if error == nil {
				response.HwReplacement = hw_replacement
			}
		}
	}

	return response, error
}

func (hwReplacementService *hwReplacementService) GetHwReplacement(request *model.GetHwReplacementRequest) (model.GetHwReplacementResponse, error) {

	var response model.GetHwReplacementResponse
	var hw_replacement []entity.HwReplacement

	ticket, error := hwReplacementService.ticketRepository.CheckTicketCode(request.TicketCode)

	if error == nil && len(ticket) > 0 {
		response.TicketCode = request.TicketCode
		response.NoSPM = ticket[0].NoSPM
		response.NoReqSPM = ticket[0].NoReqSPM

		hw_replacement, error = hwReplacementService.hwReplacementRepository.GetHwReplacement(request)

		url := os.Getenv("FILE_URL")

		for index := range hw_replacement {
			date := hw_replacement[index].CreatedAt.Format("2006-01-02")
			file_name := hw_replacement[index].Attachment
			hw_replacement[index].Attachment = url + "hw_replacement/" + request.TicketCode + "/" + date + "/" + file_name
		}

		if error == nil {
			response.HwReplacement = hw_replacement
		}
	}

	return response, error
}
