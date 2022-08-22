package service

import (
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/model"
	"svc-monitoring-maintenance/repository"
)

type ReportServiceInterface interface {
	GetReportCorrective(request *model.GetReportRequest) ([]model.GetReportCorrectiveResponse, error)
	GetReportPreventive(request *model.GetReportRequest) ([]entity.Preventive, error)
}

type reportService struct {
	repository repository.ReportRepositoryInterface
}

func ReportService(repository repository.ReportRepositoryInterface) *reportService {
	return &reportService{repository}
}

func (reportService *reportService) GetReportCorrective(request *model.GetReportRequest) ([]model.GetReportCorrectiveResponse, error) {

	request.EndDate = request.EndDate + " 23:59:59"
	ticket, error := reportService.repository.GetReportCorrective(request)

	return ticket, error

}

func (reportService *reportService) GetReportPreventive(request *model.GetReportRequest) ([]entity.Preventive, error) {

	request.EndDate = request.EndDate + " 23:59:59"
	ticket, error := reportService.repository.GetReportPreventive(request)

	return ticket, error

}
