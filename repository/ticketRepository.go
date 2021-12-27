package repository

import (
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/model"
)

type TicketRepositoryInterface interface {
	GetAll() ([]entity.Ticket, error)
	GetTicketWithoutCategory(request model.GetTicketRequest) ([]model.GetTicketResponse, error)
	GetTicketWithCategory(request model.GetTicketRequest) ([]model.GetTicketResponse, error)
	CountTicketByStatus(request model.CountTicketByStatusRequest) ([]model.CountTicketByStatusResponse, error)
	CreateTicket(request entity.Ticket) (entity.Ticket, error)
	CreateTicketIsi(request entity.TicketIsi) (entity.TicketIsi, error)
	GetTicketIsi(request string) ([]entity.TicketIsi, error)
	AssignTicket(request model.AssignTicketRequest) (entity.Ticket, error)
	UpdateTicketStatus(request model.UpdateTicketStatusRequest) (entity.Ticket, error)
	CheckTicketCode(request string) ([]model.GetTicketResponse, error)
}

func (repo *repository) GetAll() ([]entity.Ticket, error) {
	var ticket []entity.Ticket

	error := repo.db.Raw("SELECT * FROM ticket ORDER BY tgl_dibuat ASC").Scan(&ticket).Error
	return ticket, error
}

func (repo *repository) CountTicketByStatus(request model.CountTicketByStatusRequest) ([]model.CountTicketByStatusResponse, error) {
	var status []model.CountTicketByStatusResponse

	error := repo.db.Raw("SELECT status, COUNT(*) as total FROM ticket WHERE assigned_to LIKE @AssignedTo AND assigned_to_team LIKE @AssignedToTeam AND tgl_dibuat >= @StartDate AND tgl_dibuat <= @EndDate GROUP BY status", model.CountTicketByStatusRequest{
		AssignedTo:     "%" + request.AssignedTo + "%",
		AssignedToTeam: "%" + request.AssignedToTeam + "%",
		StartDate:      request.StartDate,
		EndDate:        request.EndDate,
	}).Find(&status).Error

	return status, error
}

func (repo *repository) GetTicketWithoutCategory(request model.GetTicketRequest) ([]model.GetTicketResponse, error) {
	var ticket []model.GetTicketResponse

	error := repo.db.Raw("SELECT * FROM (SELECT ticket.*, users.name as user_name, team.name as team_name, category.name as category_name FROM ticket LEFT OUTER JOIN users ON (ticket.assigned_to = CAST(users.id AS varchar(10))) LEFT OUTER JOIN team ON (ticket.assigned_to_team = CAST(team.id AS varchar(10))) LEFT OUTER JOIN category ON (ticket.category = CAST(category.id AS varchar(10))) WHERE prioritas LIKE @Priority AND status LIKE @Status AND assigned_to LIKE @AssignedTo AND assigned_to_team LIKE @AssignedToTeam AND username_pembuat LIKE @UsernamePembuat AND tgl_dibuat >= @StartDate AND tgl_dibuat <= @EndDate ORDER BY tgl_diperbarui DESC) as tbl WHERE judul LIKE @Search OR ticket_code LIKE @Search OR lokasi LIKE @Search OR terminal_id LIKE @Search OR email LIKE @Search", model.GetTicketRequest{
		Search:          "%" + request.Search + "%",
		Status:          "%" + request.Status + "%",
		Priority:        "%" + request.Priority + "%",
		AssignedTo:      "%" + request.AssignedTo + "%",
		AssignedToTeam:  "%" + request.AssignedToTeam + "%",
		UsernamePembuat: "%" + request.UsernamePembuat + "%",
		StartDate:       request.StartDate,
		EndDate:         request.EndDate,
	}).Find(&ticket).Error

	return ticket, error
}

func (repo *repository) GetTicketWithCategory(request model.GetTicketRequest) ([]model.GetTicketResponse, error) {
	var ticket []model.GetTicketResponse

	error := repo.db.Raw("SELECT * FROM (SELECT ticket.*, users.name as user_name, team.name as team_name, category.name as category_name FROM ticket LEFT OUTER JOIN users ON (ticket.assigned_to = CAST(users.id AS varchar(10))) LEFT OUTER JOIN team ON (ticket.assigned_to_team = CAST(team.id AS varchar(10))) LEFT OUTER JOIN category ON (ticket.category = CAST(category.id AS varchar(10))) WHERE ticket.category IN @Category AND prioritas LIKE @Priority AND status LIKE @Status AND assigned_to LIKE @AssignedTo AND assigned_to_team LIKE @AssignedToTeam AND username_pembuat LIKE @UsernamePembuat AND tgl_dibuat >= @StartDate AND tgl_dibuat <= @EndDate ORDER BY tgl_diperbarui DESC) as tbl WHERE judul LIKE @Search OR ticket_code LIKE @Search OR lokasi LIKE @Search OR terminal_id LIKE @Search OR email LIKE @Search", model.GetTicketRequest{
		Search:          "%" + request.Search + "%",
		Status:          "%" + request.Status + "%",
		Priority:        "%" + request.Priority + "%",
		AssignedTo:      "%" + request.AssignedTo + "%",
		AssignedToTeam:  "%" + request.AssignedToTeam + "%",
		UsernamePembuat: "%" + request.UsernamePembuat + "%",
		Category:        request.Category,
		StartDate:       request.StartDate,
		EndDate:         request.EndDate,
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

func (repo *repository) GetTicketIsi(request string) ([]entity.TicketIsi, error) {
	var ticket_isi []entity.TicketIsi

	error := repo.db.Raw("SELECT * FROM ticket_isi WHERE ticket_code = @TicketCode", entity.TicketIsi{
		TicketCode: request,
	}).Find(&ticket_isi).Error

	return ticket_isi, error
}

func (repo *repository) AssignTicket(request model.AssignTicketRequest) (entity.Ticket, error) {
	var ticket entity.Ticket

	error := repo.db.Raw("UPDATE ticket SET assigned_to = @UserId, assigned_to_team = @TeamId, tgl_diperbarui = @UpdateAt WHERE ticket_code = @TicketCode RETURNING ticket.*", model.AssignTicketRequest{
		TicketCode: request.TicketCode,
		UserId:     request.UserId,
		TeamId:     request.TeamId,
		UpdateAt:   request.UpdateAt,
	}).Find(&ticket).Error

	return ticket, error
}

func (repo *repository) UpdateTicketStatus(request model.UpdateTicketStatusRequest) (entity.Ticket, error) {
	var ticket entity.Ticket

	error := repo.db.Raw("UPDATE ticket SET status = @Status, tgl_diperbarui = @UpdateAt WHERE ticket_code = @TicketCode RETURNING ticket.*", model.UpdateTicketStatusRequest{
		TicketCode: request.TicketCode,
		Status:     request.Status,
		UpdateAt:   request.UpdateAt,
	}).Find(&ticket).Error

	return ticket, error
}

func (repo *repository) CheckTicketCode(request string) ([]model.GetTicketResponse, error) {
	var ticket []model.GetTicketResponse

	error := repo.db.Raw("SELECT ticket.*, users.name as user_name, team.name as team_name, category.name as category_name FROM ticket LEFT OUTER JOIN users ON (ticket.assigned_to = CAST(users.id AS varchar(10))) LEFT OUTER JOIN team ON (ticket.assigned_to_team = CAST(team.id AS varchar(10))) LEFT OUTER JOIN category ON (ticket.category = CAST(category.id AS varchar(10))) WHERE ticket_code = @TicketCode", model.CreateTicketRequest{
		TicketCode: request,
	}).Find(&ticket).Error

	return ticket, error
}
