package service

import (
	"fmt"
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/model"
	"svc-monitoring-maintenance/repository"
	"time"
)

type TicketServiceInterface interface {
	GetAll()													([]entity.Ticket, error)
	GetTicket(request model.GetTicketRequest)					([]model.GetTicketResponse, error)
	CountTicketByStatus(request model.CountTicketByStatusRequest)	([]model.CountTicketByStatusResponse, error)
	CreateTicket(request model.CreateTicketRequest)				(model.CreateTicketRequest, error)
	AssignTicket(request model.AssignTicketRequest) 			(entity.Ticket, error)
	UpdateTicketStatus(request model.UpdateTicketStatusRequest) (entity.Ticket, error)
	GetDetailTicket(request string) 							([]model.GetTicketResponse, error)
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

func (ticketService *ticketService) GetTicket(request model.GetTicketRequest) ([]model.GetTicketResponse, error) {
	request.EndDate = request.EndDate + " 23:59:59"
	list_ticket, error := ticketService.repository.GetTicket(request)

	return list_ticket, error
}

func (ticketService *ticketService) CountTicketByStatus(request model.CountTicketByStatusRequest) ([]model.CountTicketByStatusResponse, error){
	request.EndDate = request.EndDate + " 23:59:59"
	count_ticket, error := ticketService.repository.CountTicketByStatus(request)
	
	return count_ticket, error
}

type errorStruct struct {
	errorMessage error
}
func (ticketService *ticketService) CreateTicket(request model.CreateTicketRequest) (model.CreateTicketRequest, error) {
	date_now := time.Now()

	fmt.Printf(request.Category)
	ticket_request := entity.Ticket{
		Judul: request.Judul,
		UsernamePembuat: request.UserPembuat,
		UsernamePembalas: request.UserPembuat,
		Prioritas: request.Prioritas,
		TotalWaktu: request.TotalWaktu,
		Status: request.Status,
		TicketCode: request.TicketCode,
		Category: request.Category,
		Lokasi: request.Lokasi,
		TerminalId: request.TerminalId,
		Email: request.Email,
		AssignedTo: request.AssignedTo,
		AssignedToTeam: request.AssignedToTeam,
		TglDibuat: date_now,
		TglDiperbarui: date_now,
	}

	ticket_isi_request := entity.TicketIsi{
		UsernamePengirim: request.UserPembuat,
		Isi: request.Isi,
		TicketCode: request.TicketCode,
		Attachment1: "-",
		Attachment2: "-",
		TglDibuat: date_now,
	}

	ticket, error := ticketService.repository.CheckTicketCode(request.TicketCode)

	if (len(ticket) > 0) {
		error = fmt.Errorf("Ticket code already exist!")
	} else if (error == nil) {

		_, error = ticketService.repository.CreateTicket(ticket_request)

		if (error == nil) {
			_, error = ticketService.repository.CreateTicketIsi(ticket_isi_request)
		}
	}
	
	return request, error
	
}

func (ticketService *ticketService) AssignTicket(request model.AssignTicketRequest) (entity.Ticket, error) {

	date_now := time.Now()
	
	request.UpdateAt = date_now

	ticket, error := ticketService.repository.AssignTicket(request)

	return ticket, error
}

func (ticketService *ticketService) UpdateTicketStatus(request model.UpdateTicketStatusRequest) (entity.Ticket, error) {

	date_now := time.Now()

	request.UpdateAt = date_now

	ticket, error := ticketService.repository.UpdateTicketStatus(request)

	return ticket, error
}

func (ticketService *ticketService) GetDetailTicket(request string) ([]model.GetTicketResponse, error) {

	ticket, error := ticketService.repository.CheckTicketCode(request)
	
	ticket_isi, ticket_isi_error := ticketService.repository.GetTicketIsi(ticket[0].TicketCode)
	if (ticket_isi_error != nil) {
		error = ticket_isi_error
	}
	
	ticket[0].TicketIsi = ticket_isi

	return ticket, error
}