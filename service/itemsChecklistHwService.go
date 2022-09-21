package service

import (
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/repository"
	"time"
)

type ItemsChecklistHwServiceInterface interface {
	CreateItemsChecklistHw(request *entity.ItemsChecklistHw) ([]entity.ItemsChecklistHw, error)
	GetItemsChecklistHw() ([]entity.ItemsChecklistHw, error)
}

type itemsChecklistHwService struct {
	itemsChecklistHwRepository repository.ItemsChecklistHwRepositoryInterface
}

func ItemsChecklistHwService(itemsChecklistHwServiceInterface repository.ItemsChecklistHwRepositoryInterface) *itemsChecklistHwService {
	return &itemsChecklistHwService{itemsChecklistHwServiceInterface}
}

func (itemsChecklistHwService *itemsChecklistHwService) CreateItemsChecklistHw(request *entity.ItemsChecklistHw) ([]entity.ItemsChecklistHw, error) {

	request.CreatedAt = time.Now()
	items_checklist_hw, error := itemsChecklistHwService.itemsChecklistHwRepository.CreateItemsChecklistHw(request)

	return items_checklist_hw, error
}

func (itemsChecklistHwService *itemsChecklistHwService) GetItemsChecklistHw() ([]entity.ItemsChecklistHw, error) {

	items_checklist_hw, error := itemsChecklistHwService.itemsChecklistHwRepository.GetItemsChecklistHw()

	return items_checklist_hw, error
}
