package service

import (
	"fmt"
	"math"
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/general"
	"svc-monitoring-maintenance/model"
	"svc-monitoring-maintenance/repository"
	"time"
)

type PreventiveServiceInterface interface {
	CreatePreventive(request []entity.Preventive) ([]entity.Preventive, error)
	GetPreventive(request *model.GetPreventiveRequest) ([]model.GetGroupPreventiveResponse, int, error)
	UpdatePreventive(request *entity.Preventive) (entity.Preventive, error)
	GetDetailPreventive(request string) ([]entity.Preventive, error)
	CountPreventiveByStatus(request model.CountPreventiveByStatusRequest) ([]model.CountPreventiveByStatusResponse, error)
}

type preventiveService struct {
	repository repository.PreventiveRepositoryInterface
}

func PreventiveService(repository repository.PreventiveRepositoryInterface) *preventiveService {
	return &preventiveService{repository}
}

func (preventiveService *preventiveService) CreatePreventive(request []entity.Preventive) ([]entity.Preventive, error) {
	date_now := time.Now()
	result := []entity.Preventive{}
	var error error

	for _, value := range request {

		prev_code := "PREV-" + date_now.Format("020106") + "-" + general.RandomString(4)
		preventive_request := entity.Preventive{
			PrevCode:       prev_code,
			VisitDate:      value.VisitDate,
			Location:       value.Location,
			TerminalId:     value.TerminalId,
			AssignedTo:     value.AssignedTo,
			AssignedToTeam: value.AssignedToTeam,
			Status:         value.Status,
			CreatedBy:      value.CreatedBy,
			CreatedAt:      date_now,
			UpdatedBy:      value.CreatedBy,
			UpdatedAt:      date_now,
			AreaCode:       value.AreaCode,
			Regional:       value.Regional,
			GrapariId:      value.GrapariId,
			NoSPM:          value.NoSPM,
			NoReqSPM:       value.NoReqSPM,
			Judul:          "Preventive",
			Email:          value.Email,
		}
		_, error = preventiveService.repository.CreatePreventive(preventive_request)

		result = append(result, preventive_request)
	}

	return result, error
}

func (preventiveService *preventiveService) GetPreventive(request *model.GetPreventiveRequest) ([]model.GetGroupPreventiveResponse, int, error) {
	var preventive []entity.Preventive
	var list_group_preventive []model.GetGroupPreventiveResponse
	error := fmt.Errorf("error")
	request.EndDate = request.EndDate + " 23:59:59"
	if request.PageSize == 0 {
		request.PageSize = math.MaxInt16
	}

	request.StartIndex = request.PageNo * request.PageSize
	total_data, error := preventiveService.repository.CountVisitDate(request)
	total_pages := math.Ceil(float64(total_data) / float64(request.PageSize))
	parse_tp := int(total_pages)

	list_visit_date, error := preventiveService.repository.GetVisitDate(request)

	for _, value := range list_visit_date {

		request.StartDate = value.VisitDate
		request.EndDate = value.VisitDate + " 23:59:59"

		preventive, error = preventiveService.repository.GetPreventive(request)

		group_preventive := model.GetGroupPreventiveResponse{
			VisitDate:       value.VisitDate,
			TotalPreventive: value.TotalPreventive,
			PreventiveList:  preventive,
		}

		list_group_preventive = append(list_group_preventive, group_preventive)

	}

	return list_group_preventive, parse_tp, error

}

func (preventiveService *preventiveService) UpdatePreventive(request *entity.Preventive) (entity.Preventive, error) {
	date_now := time.Now()

	request.UpdatedAt = date_now

	preventive, error := preventiveService.repository.UpdatePreventive(request)

	return preventive, error
}

func (preventiveService *preventiveService) GetDetailPreventive(request string) ([]entity.Preventive, error) {

	preventive, error := preventiveService.repository.GetDetailPreventive(request)

	return preventive, error
}

func (preventiveService *preventiveService) CountPreventiveByStatus(request model.CountPreventiveByStatusRequest) ([]model.CountPreventiveByStatusResponse, error) {

	request.EndDate = request.EndDate + " 23:59:59"
	count_preventive, error := preventiveService.repository.CountPreventiveByStatus(request)

	return count_preventive, error
}
