package service

import (
	"svc-ticket-monitoring/entity"
	"svc-ticket-monitoring/model"
	"svc-ticket-monitoring/repository"
)

type TicketServiceInterface interface {
	FindAll()									([]entity.Ticket, error)
	FindTicket(request model.GetTicketRequest)	([]entity.Ticket, error)
}

type ticketService struct {
	repository repository.TicketRepositoryInterface
}

func TicketService(repository repository.TicketRepositoryInterface) *ticketService {
	return &ticketService{repository}
}

func (ticketService *ticketService) FindAll() ([]entity.Ticket, error){
	return ticketService.repository.FindAll()
}

func (ticketService *ticketService) FindTicket(request model.GetTicketRequest) ([]entity.Ticket, error) {
	list_ticket, error := ticketService.repository.FindTicket(request)

	return list_ticket, error
}