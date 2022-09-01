package repository

import (
	"fmt"
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/model"
)

type TicketRepositoryInterface interface {
	GetAll() ([]entity.Ticket, error)
	GetTicket(request *model.GetTicketRequest) ([]entity.Ticket, error)
	CountTicket(request *model.GetTicketRequest) (int, error)
	CountTicketByStatus(request model.CountTicketByStatusRequest) ([]model.CountTicketByStatusResponse, error)
	CreateTicket(request *entity.Ticket) (entity.Ticket, error)
	CreateTicketIsi(request *entity.TicketIsi) (entity.TicketIsi, error)
	GetTicketIsi(request string) ([]entity.TicketIsi, error)
	AssignTicket(request model.AssignTicketRequest) (entity.Ticket, error)
	UpdateTicketStatus(request model.UpdateTicketStatusRequest) (entity.Ticket, error)
	CheckTicketCode(request string) ([]model.GetTicketResponse, error)
	GetEmailHistory(request model.GetEmailHistoryRequest) ([]model.GetEmailHistoryResponse, error)
	UpdateTicket(request *entity.Ticket) (entity.Ticket, error)
	UpdateVisitStatus(request *model.UpdateVisitStatusRequest) (entity.Ticket, error)
}

func (repo *repository) GetAll() ([]entity.Ticket, error) {
	var ticket []entity.Ticket

	error := repo.db.Raw("SELECT * FROM ticket ORDER BY tgl_dibuat ASC").Scan(&ticket).Error
	return ticket, error
}

func (repo *repository) GetTicket(request *model.GetTicketRequest) ([]entity.Ticket, error) {
	var ticket []entity.Ticket
	var query string
	var category string
	var area_code string
	var regional string
	var grapari_id string

	if len(request.Category) > 0 {
		category = "AND category IN @Category"
	}
	if len(request.AreaCode) > 0 {
		area_code = "AND ticket.area_code IN @AreaCode"
	}
	if len(request.Regional) > 0 {
		regional = "AND ticket.regional IN @Regional"
	}
	if len(request.GrapariId) > 0 {
		grapari_id = "AND ticket.grapari_id IN @GrapariId"
	}

	query = fmt.Sprintf("SELECT * FROM (SELECT ticket.*, team.name as team_name, mms_category.name AS category_name, ms_area.area_name, ms_grapari.name AS grapari_name, users1.name AS user_pembuat, users2.name AS assignee FROM ticket LEFT OUTER JOIN mms_category ON (ticket.category = CAST(mms_category.id AS varchar(10))) LEFT OUTER JOIN team ON (ticket.assigned_to_team = CAST(team.id AS varchar(10))) LEFT OUTER JOIN ms_area ON (ticket.area_code = ms_area.area_code) LEFT OUTER JOIN ms_grapari ON (ticket.grapari_id = ms_grapari.grapari_id) LEFT OUTER JOIN users users1 ON (ticket.username_pembuat = CAST(users1.id AS varchar(10))) LEFT OUTER JOIN users users2 ON (ticket.assigned_to = CAST(users2.id AS varchar(10))) WHERE prioritas LIKE @Priority AND ticket.status LIKE @Status AND assigned_to LIKE @AssignedTo AND username_pembuat LIKE @UsernamePembuat %s %s %s %s AND tgl_dibuat >= @StartDate AND tgl_dibuat <= @EndDate ORDER BY tgl_diperbarui DESC) as tbl WHERE judul LIKE @Search OR ticket_code LIKE @Search OR lokasi LIKE @Search OR terminal_id LIKE @Search OR email LIKE @Search LIMIT @PageSize OFFSET @StartIndex", category, area_code, regional, grapari_id)

	error := repo.db.Raw(query, model.GetTicketRequest{
		AssignedTo:      "%" + request.AssignedTo + "%",
		Category:        request.Category,
		AreaCode:        request.AreaCode,
		Regional:        request.Regional,
		GrapariId:       request.GrapariId,
		Priority:        "%" + request.Priority + "%",
		Search:          "%" + request.Search + "%",
		Status:          "%" + request.Status + "%",
		UsernamePembuat: "%" + request.UsernamePembuat + "%",
		StartIndex:      request.StartIndex,
		PageSize:        request.PageSize,
		StartDate:       request.StartDate,
		EndDate:         request.EndDate,
	}).Find(&ticket).Error

	return ticket, error
}

func (repo *repository) CountTicket(request *model.GetTicketRequest) (int, error) {
	var total_data int
	var query string
	var category string
	var area_code string
	var regional string
	var grapari_id string

	if len(request.Category) > 0 {
		category = "AND category IN @Category"
	}
	if len(request.AreaCode) > 0 {
		area_code = "AND ticket.area_code IN @AreaCode"
	}
	if len(request.Regional) > 0 {
		regional = "AND ticket.regional IN @Regional"
	}
	if len(request.GrapariId) > 0 {
		grapari_id = "AND ticket.grapari_id IN @GrapariId"
	}

	query = fmt.Sprintf("SELECT COUNT(*) as total_data FROM (SELECT * FROM ticket WHERE prioritas LIKE @Priority AND status LIKE @Status AND assigned_to LIKE @AssignedTo AND username_pembuat LIKE @UsernamePembuat %s %s %s %s AND tgl_dibuat >= @StartDate AND tgl_dibuat <= @EndDate ORDER BY tgl_diperbarui DESC) as tbl WHERE judul LIKE @Search OR ticket_code LIKE @Search OR lokasi LIKE @Search OR terminal_id LIKE @Search OR email LIKE @Search", category, area_code, regional, grapari_id)

	error := repo.db.Raw(query, model.GetTicketRequest{
		AssignedTo:      "%" + request.AssignedTo + "%",
		Category:        request.Category,
		AreaCode:        request.AreaCode,
		Regional:        request.Regional,
		GrapariId:       request.GrapariId,
		Priority:        "%" + request.Priority + "%",
		Search:          "%" + request.Search + "%",
		Status:          "%" + request.Status + "%",
		UsernamePembuat: "%" + request.UsernamePembuat + "%",
		StartIndex:      request.StartIndex,
		PageSize:        request.PageSize,
		StartDate:       request.StartDate,
		EndDate:         request.EndDate,
	}).Find(&total_data).Error

	return total_data, error
}

func (repo *repository) CountTicketByStatus(request model.CountTicketByStatusRequest) ([]model.CountTicketByStatusResponse, error) {
	var status []model.CountTicketByStatusResponse

	var area_code string
	var regional string
	var grapari_id string

	if len(request.AreaCode) > 0 {
		area_code = "AND ticket.area_code IN @AreaCode"
	}
	if len(request.Regional) > 0 {
		regional = "AND ticket.regional IN @Regional"
	}
	if len(request.GrapariId) > 0 {
		grapari_id = "AND ticket.grapari_id IN @GrapariId"
	}

	query := fmt.Sprintf("SELECT status, COUNT(*) as total FROM ticket WHERE assigned_to LIKE @AssignedTo AND assigned_to_team LIKE @AssignedToTeam %s %s %s AND tgl_dibuat >= @StartDate AND tgl_dibuat <= @EndDate GROUP BY status", area_code, regional, grapari_id)

	error := repo.db.Raw(query, model.CountTicketByStatusRequest{
		AreaCode:       request.AreaCode,
		Regional:       request.Regional,
		GrapariId:      request.GrapariId,
		AssignedTo:     "%" + request.AssignedTo + "%",
		AssignedToTeam: "%" + request.AssignedToTeam + "%",
		StartDate:      request.StartDate,
		EndDate:        request.EndDate,
	}).Find(&status).Error

	return status, error
}

func (repo *repository) CreateTicket(request *entity.Ticket) (entity.Ticket, error) {
	var ticket entity.Ticket

	error := repo.db.Table("ticket").Create(&request).Error

	return ticket, error
}

func (repo *repository) CreateTicketIsi(request *entity.TicketIsi) (entity.TicketIsi, error) {
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

	error := repo.db.Raw("UPDATE ticket SET assigned_to = @UserId, assigned_to_team = @TeamId, tgl_diperbarui = @UpdateAt, assigning_time = @AssigningTime, assigning_by = @AssigningBy WHERE ticket_code = @TicketCode RETURNING ticket.*", request).Find(&ticket).Error

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

	error := repo.db.Raw("SELECT ticket.*, users.name as user_name, team.name as team_name, mms_category.name as category_name, ms_area.area_name, ms_grapari.name AS grapari_name, users2.name AS assignee FROM ticket LEFT OUTER JOIN users ON (ticket.assigned_to = CAST(users.id AS varchar(10))) LEFT OUTER JOIN team ON (ticket.assigned_to_team = CAST(team.id AS varchar(10))) LEFT OUTER JOIN mms_category ON (ticket.category = CAST(mms_category.id AS varchar(10))) LEFT OUTER JOIN ms_area ON (ticket.area_code = ms_area.area_code) LEFT OUTER JOIN ms_grapari ON (ticket.grapari_id = ms_grapari.grapari_id) LEFT OUTER JOIN users users2 ON (ticket.assigned_to = CAST(users2.id AS varchar(10))) WHERE ticket_code = @TicketCode", model.CreateTicketRequest{
		TicketCode: request,
	}).Find(&ticket).Error

	return ticket, error
}

func (repo *repository) GetEmailHistory(request model.GetEmailHistoryRequest) ([]model.GetEmailHistoryResponse, error) {
	var email []model.GetEmailHistoryResponse

	error := repo.db.Raw("SELECT DISTINCT ticket.email FROM ticket WHERE email LIKE @Search ORDER BY email ASC", model.GetEmailHistoryRequest{
		Search: "%" + request.Search + "%",
	}).Scan(&email).Error

	return email, error
}

func (repo *repository) UpdateTicket(request *entity.Ticket) (entity.Ticket, error) {
	var ticket entity.Ticket

	error := repo.db.Raw("UPDATE ticket SET no_spm = @NoSPM, no_req_spm = @NoReqSPM, tgl_diperbarui = @TglDiperbarui WHERE ticket_code = @TicketCode RETURNING ticket.*", request).Find(&ticket).Error

	return ticket, error
}

func (repo *repository) UpdateVisitStatus(request *model.UpdateVisitStatusRequest) (entity.Ticket, error) {
	var ticket entity.Ticket

	error := repo.db.Raw("UPDATE ticket SET visit_status = @VisitStatus, tgl_diperbarui = @UpdateAt WHERE ticket_code = @TicketCode RETURNING ticket.*", model.UpdateVisitStatusRequest{
		TicketCode:  request.TicketCode,
		VisitStatus: request.VisitStatus,
		UpdateAt:    request.UpdateAt,
	}).Find(&ticket).Error

	return ticket, error
}
