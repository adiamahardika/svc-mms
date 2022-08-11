package service

import (
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/model"
	"svc-monitoring-maintenance/repository"
)

type HardwareServiceInterface interface {
	GetHardware(request *model.GetHardwareRequest) ([]entity.Hardware, error)
}

type hardwareService struct {
	hardwareRepository repository.HardwareRepositoryInterface
}

func HardwareService(hardwareRepository repository.HardwareRepositoryInterface) *hardwareService {
	return &hardwareService{hardwareRepository}
}

func (hardwareService *hardwareService) GetHardware(request *model.GetHardwareRequest) ([]entity.Hardware, error) {

	hardware, error := hardwareService.hardwareRepository.GetHardware(request)

	return hardware, error
}
