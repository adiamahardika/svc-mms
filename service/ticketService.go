package service

import (
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/model"
	"svc-monitoring-maintenance/repository"
	"time"
)

type TicketServiceInterface interface {
	GetAll()									([]entity.Ticket, error)
	GetTicket(request model.GetTicketRequest)	([]model.GetTicketResponse, error)
	CountTicketByStatus()						([]model.CountTicketByStatusResponse, error)
	CreateTicket(request model.CreateTicketRequest)			(model.CreateTicketRequest, error)
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
		KodeTicket: request.TicketCode,
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
		KodeTicket: request.TicketCode,
		Attachment1: "-",
		UrlAttachment1: "-",
		Attachment2: "-",
		UrlAttachment2: "-",
		TglDibuat: date_now,
	}
	
	_, ticket_error := ticketService.repository.CreateTicket(ticket_request)

	
	_, ticket_isi_error := ticketService.repository.CreateTicketIsi(ticket_isi_request)
	

	if (ticket_error != nil) {
		return request, ticket_error
	} else {
		return request, ticket_isi_error
	}
}