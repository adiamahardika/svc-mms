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
	old_attachment := ""
	new_attachment := ""

	_, check_dir_error := os.Stat(path)

	if os.IsNotExist(check_dir_error) {
		check_dir_error := os.MkdirAll(path, 0755)

		if check_dir_error != nil {
			error = check_dir_error
		}
	}

	if request.OldAttachment != nil {
		old_attachment = request.OldAttachment.Filename
		error = context.SaveUploadedFile(request.OldAttachment, path+"/"+old_attachment)
	} else {
		error = nil
	}
	if request.NewAttachment != nil {
		new_attachment = request.NewAttachment.Filename
		error = context.SaveUploadedFile(request.NewAttachment, path+"/"+new_attachment)
	} else {
		error = nil
	}

	if error == nil {
		new_request := &entity.HwReplacement{
			TicketCode:    request.TicketCode,
			HwId:          request.HwId,
			OldSN:         request.OldSN,
			NewSN:         request.NewSN,
			Description:   request.Description,
			OldAttachment: old_attachment,
			NewAttachment: new_attachment,
			CreatedBy:     request.CreatedBy,
			CreatedAt:     date_now,
			StatusId:      request.StatusId,
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
				path := url + "hw_replacement/" + request.TicketCode + "/" + date + "/"
				if hw_replacement[index].OldAttachment != "" {
					file_name1 := hw_replacement[index].OldAttachment
					hw_replacement[index].OldAttachment = path + file_name1
				}
				if hw_replacement[index].NewAttachment != "" {
					file_name2 := hw_replacement[index].NewAttachment
					hw_replacement[index].NewAttachment = path + file_name2
				}
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
			path := url + "hw_replacement/" + request.TicketCode + "/" + date + "/"

			if hw_replacement[index].OldAttachment != "" {
				file_name1 := hw_replacement[index].OldAttachment
				hw_replacement[index].OldAttachment = path + file_name1
			}
			if hw_replacement[index].NewAttachment != "" {
				file_name2 := hw_replacement[index].NewAttachment
				hw_replacement[index].NewAttachment = path + file_name2
			}
		}

		if error == nil {
			response.HwReplacement = hw_replacement
		}
	}

	return response, error
}
