package service

import (
	"svc-ticket-monitoring/entity"
	"svc-ticket-monitoring/model"
	"svc-ticket-monitoring/repository"
)

type TicketServiceInterface interface {
	GetAll()									([]entity.Ticket, error)
	GetTicket(request model.GetTicketRequest)	([]entity.Ticket, error)
	CountTicketByStatus()						([]model.CountTicketByStatusResponse, error)
}

type ticketService struct {
	repository repository.TicketRepositoryInterface
}

func TicketService(repository repository.TicketRepositoryInterface) *ticketService {
	return &ticketService{repository}
}

func (ticketService *ticketService) GetAll() ([]entity.Ticket, error){
	return ticketService.repository.GetAll()
}

func (ticketService *ticketService) GetTicket(request model.GetTicketRequest) ([]entity.Ticket, error) {
	list_ticket, error := ticketService.repository.GetTicket(request)

	return list_ticket, error
}

func (ticketService *ticketService) CountTicketByStatus() ([]model.CountTicketByStatusResponse, error){
	return ticketService.repository.CountTicketByStatus()
}