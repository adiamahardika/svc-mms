package repository

import (
	"svc-ticket-monitoring/entity"
	"svc-ticket-monitoring/model"
)

type TicketRepositoryInterface interface {
	GetAll() 										([]entity.Ticket, error)
	GetTicket(request model.GetTicketRequest)		([]entity.Ticket, error)
	CountTicketByStatus()							([]model.CountTicketByStatusResponse, error)
	CreateTicket(request entity.Ticket)				(entity.Ticket, error)
	CreateTicketIsi(request entity.TicketIsi) (entity.TicketIsi, error)
}

func (repo *repository) GetAll() ([]entity.Ticket, error) {
	var ticket []entity.Ticket

	error := repo.db.Raw("SELECT * FROM ticket ORDER BY tgl_dibuat ASC").Scan(&ticket).Error 
	return ticket, error
}

func (repo *repository) CountTicketByStatus() ([]model.CountTicketByStatusResponse, error) {
	var status []model.CountTicketByStatusResponse
	
	error := repo.db.Raw("SELECT status, COUNT(*) as total FROM ticket GROUP BY status").Find(&status).Error
	return status, error
}

func (repo *repository) GetTicket(request model.GetTicketRequest) ([]entity.Ticket, error) {
	var ticket []entity.Ticket
	
	error := repo.db.Raw("SELECT * FROM (SELECT * FROM ticket WHERE prioritas LIKE @Priority AND status LIKE @Status AND assigned_to LIKE @AssignedTo AND assigned_to_team LIKE @AssignedToTeam AND username_pembuat LIKE @UsernamePembuat AND tgl_dibuat >= @StartDate AND tgl_dibuat <= @EndDate ORDER BY tgl_diperbarui DESC) as tbl WHERE judul LIKE @Search OR kode_ticket LIKE @Search OR lokasi LIKE @Search OR terminal_id LIKE @Search OR email LIKE @Search", model.GetTicketRequest{
		PageNo:          request.PageNo,
		PageSize:        request.PageSize,
		SortBy:          "%" + request.SortBy + "%",
		Search:          "%" + request.Search + "%",
		Status:          "%" + request.Status + "%",
		Priority:        "%" + request.Priority + "%",
		AssignedTo:      "%" + request.AssignedTo + "%",
		AssignedToTeam:  "%" + request.AssignedToTeam + "%",
		UsernamePembuat: "%" + request.UsernamePembuat + "%",
		StartDate: 		request.StartDate,
		EndDate: 		request.EndDate,
		}).Find(&ticket).Error
		
		return ticket, error
}

func (repo *repository) CreateTicket(request entity.Ticket) (entity.Ticket, error) {
	var ticket entity.Ticket

	error := repo.db.Table("ticket").Create(&request).Error

	return ticket, error
}

func (repo *repository) CreateTicketIsi(request entity.TicketIsi) (entity.TicketIsi, error) {
	var ticket_isi entity.TicketIsi

	error := repo.db.Table("ticket_isi").Create(&request).Error

	return ticket_isi, error
}

