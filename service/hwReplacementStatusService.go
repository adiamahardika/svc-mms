package service

import (
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/repository"
)

type HwReplacementStatusServiceInterface interface {
	GetHwReplacementStatus() ([]entity.HwReplacementStatus, error)
	CreateHwReplacementStatus(request *entity.HwReplacementStatus) ([]entity.HwReplacementStatus, error)
	UpdateHwReplacementStatus(request *entity.HwReplacementStatus) ([]entity.HwReplacementStatus, error)
}

type hwReplacementStatusService struct {
	hwReplacementStatusRepository repository.HwReplacementStatusRepositoryInterface
}

func HwReplacementStatusService(hwReplacementStatusRepository repository.HwReplacementStatusRepositoryInterface) *hwReplacementStatusService {
	return &hwReplacementStatusService{hwReplacementStatusRepository}
}

func (hwReplacementStatusService *hwReplacementStatusService) GetHwReplacementStatus() ([]entity.HwReplacementStatus, error) {

	status, error := hwReplacementStatusService.hwReplacementStatusRepository.GetHwReplacementStatus()

	return status, error
}

func (hwReplacementStatusService *hwReplacementStatusService) CreateHwReplacementStatus(request *entity.HwReplacementStatus) ([]entity.HwReplacementStatus, error) {

	request.IsActive = "true"
	status, error := hwReplacementStatusService.hwReplacementStatusRepository.CreateHwReplacementStatus(request)

	return status, error
}

func (hwReplacementStatusService *hwReplacementStatusService) UpdateHwReplacementStatus(request *entity.HwReplacementStatus) ([]entity.HwReplacementStatus, error) {

	status, error := hwReplacementStatusService.hwReplacementStatusRepository.UpdateHwReplacementStatus(request)

	return status, error
}
