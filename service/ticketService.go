package service

import (
	"fmt"
	"math"
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/model"
	"svc-monitoring-maintenance/repository"
	"time"
)

type TicketServiceInterface interface {
	GetAll() ([]entity.Ticket, error)
	GetTicket(request *model.GetTicketRequest) ([]entity.Ticket, int, error)
	CountTicketByStatus(request model.CountTicketByStatusRequest) ([]model.CountTicketByStatusResponse, error)
	CreateTicket(request *model.CreateTicketRequest) (model.CreateTicketRequest, error)
	AssignTicket(request model.AssignTicketRequest) (entity.Ticket, error)
	UpdateTicketStatus(request model.UpdateTicketStatusRequest) (entity.Ticket, error)
	GetDetailTicket(request string) ([]model.GetTicketResponse, error)
	GetEmailHistory(request model.GetEmailHistoryRequest) ([]model.GetEmailHistoryResponse, error)
	UpdateTicket(request *entity.Ticket) (entity.Ticket, error)
	UpdateVisitStatus(request *model.UpdateVisitStatusRequest) (entity.Ticket, error)
	GetTicketActivity(request *model.GetTicketActivityRequest) ([]model.GetTicketActivityResponse, error)
}

type ticketService struct {
	repository repository.TicketRepositoryInterface
}

func TicketService(repository repository.TicketRepositoryInterface) *ticketService {
	return &ticketService{repository}
}

func (ticketService *ticketService) GetAll() ([]entity.Ticket, error) {
	return ticketService.repository.GetAll()
}

func (ticketService *ticketService) GetTicket(request *model.GetTicketRequest) ([]entity.Ticket, int, error) {

	if request.PageSize == 0 {
		request.PageSize = math.MaxInt16
	}
	request.StartIndex = request.PageNo * request.PageSize
	request.EndDate = request.EndDate + " 23:59:59"
	total_data, error := ticketService.repository.CountTicket(request)
	total_pages := math.Ceil(float64(total_data) / float64(request.PageSize))

	ticket, error := ticketService.repository.GetTicket(request)
	parse_tp := int(total_pages)

	return ticket, parse_tp, error

}

func (ticketService *ticketService) CountTicketByStatus(request model.CountTicketByStatusRequest) ([]model.CountTicketByStatusResponse, error) {
	request.EndDate = request.EndDate + " 23:59:59"
	count_ticket, error := ticketService.repository.CountTicketByStatus(request)

	return count_ticket, error
}

type errorStruct struct {
	errorMessage error
}

func (ticketService *ticketService) CreateTicket(request *model.CreateTicketRequest) (model.CreateTicketRequest, error) {
	date_now := time.Now()

	ticket_request := &entity.Ticket{
		Judul:           request.Judul,
		UsernamePembuat: request.UserPembuat,
		UpdatedBy:       request.UserPembuat,
		Prioritas:       request.Prioritas,
		Status:          request.Status,
		TicketCode:      request.TicketCode,
		Category:        request.Category,
		Lokasi:          request.Lokasi,
		TerminalId:      request.TerminalId,
		Email:           request.Email,
		AssignedTo:      request.AssignedTo,
		AssignedToTeam:  request.AssignedToTeam,
		NoSPM:           request.NoSPM,
		NoReqSPM:        request.NoReqSPM,
		TglDibuat:       date_now,
		TglDiperbarui:   date_now,
		AreaCode:        request.AreaCode,
		Regional:        request.Regional,
		GrapariId:       request.GrapariId,
		SubCategory:     request.SubCategory,
	}

	ticket_isi_request := &entity.TicketIsi{
		UsernamePengirim: request.UserPembuat,
		Isi:              request.Isi,
		TicketCode:       request.TicketCode,
		Attachment1:      "-",
		Attachment2:      "-",
		TglDibuat:        date_now,
	}

	ticket, error := ticketService.repository.CheckTicketCode(request.TicketCode)

	if len(ticket) > 0 {
		error = fmt.Errorf("Ticket code already exist!")
	} else if error == nil {

		_, error = ticketService.repository.CreateTicket(ticket_request)

		if error == nil {
			_, error = ticketService.repository.CreateTicketIsi(ticket_isi_request)
		}
	}

	return *request, error

}

func (ticketService *ticketService) AssignTicket(request model.AssignTicketRequest) (entity.Ticket, error) {

	date_now := time.Now()

	request.UpdateAt = date_now
	request.AssigningTime = date_now

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
	if ticket_isi_error != nil {
		error = ticket_isi_error
	}

	ticket[0].TicketIsi = ticket_isi

	return ticket, error
}

func (ticketService *ticketService) GetEmailHistory(request model.GetEmailHistoryRequest) ([]model.GetEmailHistoryResponse, error) {

	email, error := ticketService.repository.GetEmailHistory(request)

	return email, error
}

func (ticketService *ticketService) UpdateTicket(request *entity.Ticket) (entity.Ticket, error) {

	date_now := time.Now()

	request.TglDiperbarui = date_now

	ticket, error := ticketService.repository.UpdateTicket(request)

	return ticket, error
}

func (ticketService *ticketService) UpdateVisitStatus(request *model.UpdateVisitStatusRequest) (entity.Ticket, error) {

	date_now := time.Now()

	request.UpdateAt = date_now

	ticket, error := ticketService.repository.UpdateVisitStatus(request)

	return ticket, error
}

func (ticketService *ticketService) GetTicketActivity(request *model.GetTicketActivityRequest) ([]model.GetTicketActivityResponse, error) {

	request.EndDate = request.EndDate + " 23:59:59"
	ticket, error := ticketService.repository.GetTicketActivity(request)

	return ticket, error
}
