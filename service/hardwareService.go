package service

import (
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/model"
	"svc-monitoring-maintenance/repository"
	"time"
)

type HardwareServiceInterface interface {
	GetHardware(request *model.GetHardwareRequest) ([]entity.Hardware, error)
	CreateHardware(request *entity.Hardware) (entity.Hardware, error)
	UpdateHardware(request *entity.Hardware) (entity.Hardware, error)
	DeleteHardware(id *int) error
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

func (hardwareService *hardwareService) CreateHardware(request *entity.Hardware) (entity.Hardware, error) {

	request.CreatedAt = time.Now()
	request.IsActive = "true"

	hardware, error := hardwareService.hardwareRepository.CreateHardware(request)

	return hardware, error
}

func (hardwareService *hardwareService) UpdateHardware(request *entity.Hardware) (entity.Hardware, error) {

	hardware, error := hardwareService.hardwareRepository.UpdateHardware(request)

	return hardware, error
}

func (hardwareService *hardwareService) DeleteHardware(id *int) error {

	error := hardwareService.hardwareRepository.DeleteHardware(id)

	return error
}
