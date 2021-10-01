package service

import (
	"fmt"
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/model"
	"svc-monitoring-maintenance/repository"
	"time"
)

type TicketServiceInterface interface {
	GetAll()														([]entity.Ticket, error)
	GetTicket(request model.GetTicketRequest)						([]model.GetTicketResponse, error)
	CountTicketByStatus()											([]model.CountTicketByStatusResponse, error)
	CreateTicket(request model.CreateTicketRequest)					(model.CreateTicketRequest, error)
	AssignTicketToMember(request model.AssignTicketToMemberRequest) (entity.Ticket, error)
	UpdateTicketStatus(request model.UpdateTicketStatusRequest) (entity.Ticket, error)
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
	list_ticket, error := ticketService.repository.GetTicket(request)

	return list_ticket, error
}

func (ticketService *ticketService) CountTicketByStatus() ([]model.CountTicketByStatusResponse, error){
	return ticketService.repository.CountTicketByStatus()
}

type errorStruct struct {
	errorMessage error
}
func (ticketService *ticketService) CreateTicket(request model.CreateTicketRequest) (model.CreateTicketRequest, error) {
	date_now := time.Now()

	ticket_request := entity.Ticket{
		Judul: request.Judul,
		UsernamePembuat: request.UserPembuat,
		UsernamePembalas: request.UserPembuat,
		Prioritas: request.Prioritas,
		TotalWaktu: request.TotalWaktu,
		Status: request.Status,
		TicketCode: request.TicketCode,
		Kategori: request.Kategori,
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
		UrlAttachment1: "-",
		Attachment2: "-",
		UrlAttachment2: "-",
		TglDibuat: date_now,
	}

	ticket, error := ticketService.repository.CheckTicketCode(request.TicketCode)

	if (len(ticket) > 0) {
		error = fmt.Errorf("Ticket code already exist!")
	} else if (error != nil) {
		
		_, error = ticketService.repository.CreateTicket(ticket_request)
	
		if (error != nil) {
			_, error = ticketService.repository.CreateTicketIsi(ticket_isi_request)
		}
	}
	
	return request, error
	
}

func (ticketService *ticketService) AssignTicketToMember(request model.AssignTicketToMemberRequest) (entity.Ticket, error) {

	date_now := time.Now()
	
	request.UpdateAt = date_now

	ticket, error := ticketService.repository.AssignTicketToMember(request)

	return ticket, error
}

func (ticketService *ticketService) UpdateTicketStatus(request model.UpdateTicketStatusRequest) (entity.Ticket, error) {

	date_now := time.Now()

	request.UpdateAt = date_now

	ticket, error := ticketService.repository.UpdateTicketStatus(request)

	return ticket, error
}