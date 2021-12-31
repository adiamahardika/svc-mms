package service

import (
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/general"
	"svc-monitoring-maintenance/model"
	"svc-monitoring-maintenance/repository"
	"time"
)

type PreventiveServiceInterface interface {
	CreatePreventive(request []model.CreatePreventiveRequest) ([]entity.Preventive, error)
	GetPreventive(request model.GetPreventiveRequest) ([]model.GetPreventiveResponse, error)
}

type preventiveService struct {
	repository repository.PreventiveRepositoryInterface
}

func PreventiveService(repository repository.PreventiveRepositoryInterface) *preventiveService {
	return &preventiveService{repository}
}

func (preventiveService *preventiveService) CreatePreventive(request []model.CreatePreventiveRequest) ([]entity.Preventive, error) {
	date_now := time.Now()
	result := []entity.Preventive{}
	var error error

	for _, value := range request {

		prev_code := "PREV-" + date_now.Format("020106") + "-" + general.RandomString(4)
		preventive_request := entity.Preventive{
			PrevCode:   prev_code,
			VisitDate:  value.VisitDate,
			TerminalId: value.TerminalId,
			AssignedTo: value.AssignedTo,
			Status:     value.Status,
			CreatedBy:  value.CreatedBy,
			CreatedAt:  date_now,
			UpdatedBy:  value.CreatedBy,
			UpdatedAt:  date_now,
		}
		_, error = preventiveService.repository.CreatePreventive(preventive_request)

		result = append(result, preventive_request)
	}

	return result, error
}

func (preventiveService *preventiveService) GetPreventive(request model.GetPreventiveRequest) ([]model.GetPreventiveResponse, error) {
	request.EndDate = request.EndDate + " 23:59:59"

	preventive, error := preventiveService.repository.GetPreventive(request)

	return preventive, error

}
