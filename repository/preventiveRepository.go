package repository

import (
	"fmt"
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/model"
)

type PreventiveRepositoryInterface interface {
	CreatePreventive(request entity.Preventive) (entity.Preventive, error)
	GetPreventive(request *model.GetPreventiveRequest) ([]entity.Preventive, error)
	CountVisitDate(request *model.GetPreventiveRequest) (int, error)
	UpdatePreventive(request *entity.Preventive) (entity.Preventive, error)
	GetDetailPreventive(request string) ([]entity.Preventive, error)
	GetVisitDate(request *model.GetPreventiveRequest) ([]model.GetVisitDateResponse, error)
	CountPreventiveByStatus(request model.CountPreventiveByStatusRequest) ([]model.CountPreventiveByStatusResponse, error)
}

func (repo *repository) CreatePreventive(request entity.Preventive) (entity.Preventive, error) {
	var preventive entity.Preventive

	error := repo.db.Table("preventive").Create(&request).Error

	return preventive, error
}

func (repo *repository) GetPreventive(request *model.GetPreventiveRequest) ([]entity.Preventive, error) {
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

	query := fmt.Sprintf("SELECT * FROM (SELECT preventive.*, users.name AS user_name, team.name as team_name, ms_area.area_name, ms_grapari.name AS grapari_name, users2.name AS creator FROM preventive LEFT OUTER JOIN users ON (preventive.assigned_to = CAST(users.id AS varchar(10))) LEFT OUTER JOIN users users2 ON (preventive.created_by = CAST(users2.id AS varchar(10))) LEFT OUTER JOIN team ON (preventive.assigned_to_team = CAST(team.id AS varchar(10))) LEFT OUTER JOIN ms_area ON (preventive.area_code = ms_area.area_code) LEFT OUTER JOIN ms_grapari ON (preventive.grapari_id = ms_grapari.grapari_id) WHERE preventive.status IN @Status AND assigned_to LIKE @AssignedTo AND assigned_to_team LIKE @AssignedToTeam %s %s %s AND visit_date >= @StartDate AND visit_date <= @EndDate ORDER BY visit_date DESC) AS tbl WHERE LOWER(tbl.terminal_id) LIKE LOWER(@Search) OR LOWER(tbl.location) LIKE LOWER(@Search)", area_code, grapari_id, regional)

	error := repo.db.Raw(query, model.GetPreventiveRequest{
		Search:         "%" + request.Search + "%",
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

func (repo *repository) CountVisitDate(request *model.GetPreventiveRequest) (int, error) {
	var preventive int
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

	query := fmt.Sprintf("SELECT COUNT(*) as total_data FROM (SELECT DISTINCT ON (visit_date) preventive.* FROM preventive WHERE preventive.status IN @Status AND assigned_to LIKE @AssignedTo AND assigned_to_team LIKE @AssignedToTeam %s %s %s AND visit_date >= @StartDate AND visit_date <= @EndDate ORDER BY visit_date DESC) AS tbl WHERE LOWER(tbl.terminal_id) LIKE LOWER(@Search) OR LOWER(tbl.location) LIKE LOWER(@Search)", area_code, grapari_id, regional)

	error := repo.db.Raw(query, model.GetPreventiveRequest{
		Search:         "%" + request.Search + "%",
		Status:         request.Status,
		AssignedTo:     "%" + request.AssignedTo + "%",
		AssignedToTeam: "%" + request.AssignedToTeam + "%",
		StartDate:      request.StartDate,
		EndDate:        request.EndDate,
		AreaCode:       request.AreaCode,
		Regional:       request.Regional,
		GrapariId:      request.GrapariId,
	}).Find(&preventive).Error

	return preventive, error
}

func (repo *repository) GetVisitDate(request *model.GetPreventiveRequest) ([]model.GetVisitDateResponse, error) {
	var list_visit_date []model.GetVisitDateResponse
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

	query := fmt.Sprintf("SELECT * FROM (SELECT preventive.visit_date, COUNT(*) AS total_preventive FROM preventive WHERE preventive.status IN @Status AND assigned_to LIKE @AssignedTo AND assigned_to_team LIKE @AssignedToTeam AND LOWER(terminal_id) LIKE LOWER(@Search) %s %s %s AND visit_date >= @StartDate AND visit_date <= @EndDate GROUP BY visit_date) AS tbl ORDER BY tbl.visit_date DESC LIMIT @PageSize OFFSET @StartIndex", area_code, grapari_id, regional)

	error := repo.db.Raw(query, model.GetPreventiveRequest{
		Search:         "%" + request.Search + "%",
		Status:         request.Status,
		AssignedTo:     "%" + request.AssignedTo + "%",
		AssignedToTeam: "%" + request.AssignedToTeam + "%",
		StartDate:      request.StartDate,
		EndDate:        request.EndDate,
		StartIndex:     request.StartIndex,
		PageSize:       request.PageSize,
		AreaCode:       request.AreaCode,
		Regional:       request.Regional,
		GrapariId:      request.GrapariId,
	}).Find(&list_visit_date).Error

	return list_visit_date, error
}

func (repo *repository) UpdatePreventive(request *entity.Preventive) (entity.Preventive, error) {
	var preventive entity.Preventive

	error := repo.db.Raw("UPDATE preventive SET visit_date = @VisitDate, area_code = @AreaCode, regional = @Regional, grapari_id = @GrapariId, location = @Location, terminal_id = @TerminalId, assigned_to = @AssignedTo, assigned_to_team = @AssignedToTeam, updated_by = @UpdatedBy, updated_at = @UpdatedAt, judul = @Judul, Note = @Note, status = @Status, no_spm = @NoSPM, no_req_spm = @NoReqSPM, email = @Email WHERE prev_code = @PrevCode RETURNING preventive.*", request).Find(&preventive).Error

	return preventive, error
}

func (repo *repository) GetDetailPreventive(request string) ([]entity.Preventive, error) {
	var preventive []entity.Preventive

	error := repo.db.Raw("SELECT preventive.*, users.name as user_name, team.name as team_name, ms_area.area_name, ms_grapari.name AS grapari_name, users2.name AS creator FROM preventive LEFT OUTER JOIN users ON (preventive.assigned_to = CAST(users.id AS varchar(10))) LEFT OUTER JOIN team ON (preventive.assigned_to_team = CAST(team.id AS varchar(10))) LEFT OUTER JOIN ms_area ON (preventive.area_code = ms_area.area_code) LEFT OUTER JOIN ms_grapari ON (preventive.grapari_id = ms_grapari.grapari_id) LEFT OUTER JOIN users users2 ON (preventive.created_by = CAST(users2.id AS varchar(10))) WHERE prev_code = @PrevCode", entity.Preventive{
		PrevCode: request,
	}).Find(&preventive).Error

	return preventive, error
}

func (repo *repository) CountPreventiveByStatus(request model.CountPreventiveByStatusRequest) ([]model.CountPreventiveByStatusResponse, error) {
	var status []model.CountPreventiveByStatusResponse

	error := repo.db.Raw("SELECT status, COUNT(*) AS total FROM preventive WHERE assigned_to LIKE @AssignedTo AND assigned_to_team LIKE @AssignedToTeam AND CAST(visit_date AS DATE) >= @StartDate AND CAST(visit_date AS DATE) <= @EndDate GROUP BY status", model.CountPreventiveByStatusRequest{
		AssignedTo:     "%" + request.AssignedTo + "%",
		AssignedToTeam: "%" + request.AssignedToTeam + "%",
		StartDate:      request.StartDate,
		EndDate:        request.EndDate,
	}).Find(&status).Error

	return status, error
}
