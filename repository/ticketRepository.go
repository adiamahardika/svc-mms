package repository

import (
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/model"
)

type TicketRepositoryInterface interface {
	GetAll() 										([]entity.Ticket, error)
	GetTicket(request model.GetTicketRequest)		([]model.GetTicketResponse, error)
	CountTicketByStatus()							([]model.CountTicketByStatusResponse, error)
	CreateTicket(request entity.Ticket)				(entity.Ticket, error)
	CreateTicketIsi(request entity.TicketIsi) 		(entity.TicketIsi, error)
	AssignTicketToMember(request model.AssignTicketToMemberRequest) (entity.Ticket, error)
	UpdateTicketStatus(request model.UpdateTicketStatusRequest) (entity.Ticket, error)
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

func (repo *repository) GetTicket(request model.GetTicketRequest) ([]model.GetTicketResponse, error) {
	var ticket []model.GetTicketResponse
	
	error := repo.db.Raw("SELECT * FROM (SELECT ticket.*, users.name as user_name, team.name as team_name FROM ticket LEFT OUTER JOIN users ON (ticket.assigned_to = CAST(users.id AS varchar(10))) LEFT OUTER JOIN team ON (ticket.assigned_to_team = CAST(team.id AS varchar(10))) WHERE prioritas LIKE @Priority AND status LIKE @Status AND assigned_to LIKE @AssignedTo AND assigned_to_team LIKE @AssignedToTeam AND username_pembuat LIKE @UsernamePembuat AND tgl_dibuat >= @StartDate AND tgl_dibuat <= @EndDate ORDER BY tgl_diperbarui DESC) as tbl WHERE judul LIKE @Search OR kode_ticket LIKE @Search OR lokasi LIKE @Search OR terminal_id LIKE @Search OR email LIKE @Search", model.GetTicketRequest{
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

func (repo *repository) AssignTicketToMember(request model.AssignTicketToMemberRequest) (entity.Ticket, error) {
	var ticket entity.Ticket

	error := repo.db.Raw("UPDATE ticket SET assigned_to = @UserId WHERE id = @Id RETURNING ticket.*", model.AssignTicketToMemberRequest{
		Id: request.Id,
		UserId: request.UserId,
	}).Find(&ticket).Error

	return ticket, error
}

func (repo *repository) UpdateTicketStatus(request model.UpdateTicketStatusRequest) (entity.Ticket, error) {
	var ticket entity.Ticket

	error := repo.db.Raw("UPDATE ticket SET status = @Status WHERE id = @Id RETURNING ticket.*", model.UpdateTicketStatusRequest{
		Id: request.Id,
		Status: request.Status,
	}).Find(&ticket).Error

	return ticket, error
}