package service

import (
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/repository"
)

type HwReplacementStatusServiceInterface interface {
	GetHwReplacementStatus() ([]entity.HwReplacementStatus, error)
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
