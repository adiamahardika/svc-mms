package repository

import (
	"fmt"
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/model"
)

type ReportRepositoryInterface interface {
	GetReportCorrective(request *model.GetReportRequest) ([]model.GetReportCorrectiveResponse, error)
	GetReportPreventive(request *model.GetReportRequest) ([]entity.Preventive, error)
}

func (repo *repository) GetReportCorrective(request *model.GetReportRequest) ([]model.GetReportCorrectiveResponse, error) {
	var ticket []model.GetReportCorrectiveResponse
	var query string
	var category string
	var created_by string
	var area_code string
	var regional string
	var grapari_id string

	if len(request.Category) > 0 {
		category = "AND mms_category IN @Category"
	}
	if len(request.UsernamePembuat) > 0 {
		created_by = "AND username_pembuat IN @UsernamePembuat"
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

	query = fmt.Sprintf("SELECT DISTINCT ON (ticket_code) ticket.*, TO_CHAR(ticket.tgl_diperbarui, 'DD-MM-YYYY HH24:MI:SS') AS tgl_diperbarui, TO_CHAR(ticket.tgl_dibuat, 'DD-MM-YYYY HH24:MI:SS' ) AS tgl_dibuat, TO_CHAR(ticket.start_time, 'DD-MM-YYYY HH24:MI:SS' ) AS start_time, TO_CHAR(ticket.close_time, 'DD-MM-YYYY HH24:MI:SS' ) AS close_time, TO_CHAR(ticket.assigning_time, 'DD-MM-YYYY HH24:MI:SS' ) AS assigning_time, team.name as team_name, mms_category.name AS category_name, ms_area.area_name, ms_grapari.name AS grapari_name, users1.name AS user_pembuat, users2.name AS assignee, ticket_isi.isi FROM ticket LEFT OUTER JOIN mms_category ON (ticket.category = CAST(mms_category.id AS varchar(10))) LEFT OUTER JOIN team ON (ticket.assigned_to_team = CAST(team.id AS varchar(10))) LEFT OUTER JOIN ms_area ON (ticket.area_code = ms_area.area_code) LEFT OUTER JOIN ms_grapari ON (ticket.grapari_id = ms_grapari.grapari_id) LEFT OUTER JOIN users users1 ON (ticket.username_pembuat = CAST(users1.id AS varchar(10))) LEFT OUTER JOIN users users2 ON (ticket.assigned_to = CAST(users2.id AS varchar(10)))  LEFT OUTER JOIN ticket_isi ON (ticket.ticket_code = ticket_isi.ticket_code) WHERE prioritas IN @Priority AND ticket.status IN @Status AND assigned_to LIKE @AssignedTo AND assigned_to_team LIKE @AssignedToTeam %s %s %s %s %s AND ticket.tgl_dibuat >= @StartDate AND ticket.tgl_dibuat <= @EndDate ORDER BY ticket_code DESC", category, area_code, regional, grapari_id, created_by)

	error := repo.db.Raw(query, model.GetReportRequest{
		AssignedTo:      "%" + request.AssignedTo + "%",
		AssignedToTeam:  "%" + request.AssignedToTeam + "%",
		Category:        request.Category,
		AreaCode:        request.AreaCode,
		Regional:        request.Regional,
		GrapariId:       request.GrapariId,
		Priority:        request.Priority,
		Status:          request.Status,
		UsernamePembuat: request.UsernamePembuat,
		StartDate:       request.StartDate,
		EndDate:         request.EndDate,
	}).Find(&ticket).Error

	return ticket, error
}

func (repo *repository) GetReportPreventive(request *model.GetReportRequest) ([]entity.Preventive, error) {
	var preventive []entity.Preventive
	var area_code string
	var regional string
	var grapari_id string

	if len(request.AreaCode) > 0 {
		area_code = "AND preventive.area_code IN @AreaCode"
	}
	if len(request.Regional) > 0 {
		regional = "AND preventive.regional IN @Regional"
	}
	if len(request.GrapariId) > 0 {
		grapari_id = "AND preventive.grapari_id IN @GrapariId"
	}

	query := fmt.Sprintf("SELECT preventive.*, users.name AS user_name, team.name as team_name, ms_area.area_name, ms_grapari.name AS grapari_name, users2.name AS creator FROM preventive LEFT OUTER JOIN users ON (preventive.assigned_to = CAST(users.id AS varchar(10))) LEFT OUTER JOIN users users2 ON (preventive.created_by = CAST(users2.id AS varchar(10))) LEFT OUTER JOIN team ON (preventive.assigned_to_team = CAST(team.id AS varchar(10))) LEFT OUTER JOIN ms_area ON (preventive.area_code = ms_area.area_code) LEFT OUTER JOIN ms_grapari ON (preventive.grapari_id = ms_grapari.grapari_id) WHERE preventive.status IN @Status AND assigned_to LIKE @AssignedTo AND assigned_to_team LIKE @AssignedToTeam %s %s %s AND visit_date >= @StartDate AND visit_date <= @EndDate ORDER BY visit_date DESC", area_code, grapari_id, regional)

	error := repo.db.Raw(query, model.GetReportRequest{
		Status:         request.Status,
		AssignedTo:     "%" + request.AssignedTo + "%",
		AssignedToTeam: "%" + request.AssignedToTeam + "%",
		AreaCode:       request.AreaCode,
		Regional:       request.Regional,
		GrapariId:      request.GrapariId,
		StartDate:      request.StartDate,
		EndDate:        request.EndDate,
	}).Find(&preventive).Error

	return preventive, error
}
