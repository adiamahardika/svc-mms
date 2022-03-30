package repository

import (
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/model"
)

type PreventiveRepositoryInterface interface {
	CreatePreventive(request entity.Preventive) (entity.Preventive, error)
	GetPreventive(request model.GetPreventiveRequest) ([]model.GetPreventiveResponse, error)
	UpdatePreventive(request model.UpdatePreventiveRequest) (entity.Preventive, error)
	GetDetailPreventive(request string) ([]model.GetPreventiveResponse, error)
	GetVisitDate(request model.GetPreventiveRequest) ([]model.GetVisitDateResponse, error)
	CountPreventiveByStatus(request model.CountPreventiveByStatusRequest) ([]model.CountPreventiveByStatusResponse, error)
}

func (repo *repository) CreatePreventive(request entity.Preventive) (entity.Preventive, error) {
	var preventive entity.Preventive

	error := repo.db.Table("preventive").Create(&request).Error

	return preventive, error
}

func (repo *repository) GetPreventive(request model.GetPreventiveRequest) ([]model.GetPreventiveResponse, error) {
	var preventive []model.GetPreventiveResponse

	error := repo.db.Raw("SELECT * FROM (SELECT preventive.*, users.name AS user_name FROM preventive LEFT OUTER JOIN users ON (preventive.assigned_to = CAST(users.id AS varchar(10))) WHERE status LIKE @Status AND assigned_to LIKE @AssignedTo AND visit_date >= @StartDate AND visit_date <= @EndDate ORDER BY visit_date DESC) AS tbl WHERE tbl.terminal_id LIKE @Search OR tbl.location LIKE @Search", model.GetPreventiveRequest{
		Search:     "%" + request.Search + "%",
		Status:     "%" + request.Status + "%",
		AssignedTo: "%" + request.AssignedTo + "%",
		StartDate:  request.StartDate,
		EndDate:    request.EndDate,
	}).Find(&preventive).Error

	return preventive, error
}

func (repo *repository) GetVisitDate(request model.GetPreventiveRequest) ([]model.GetVisitDateResponse, error) {
	var list_visit_date []model.GetVisitDateResponse

	error := repo.db.Raw("SELECT * FROM (SELECT preventive.visit_date, COUNT(*) AS total_preventive FROM preventive WHERE status LIKE @Status AND assigned_to LIKE @AssignedTo AND terminal_id LIKE @Search AND visit_date >= @StartDate AND visit_date <= @EndDate GROUP BY visit_date) AS tbl ORDER BY tbl.visit_date DESC", model.GetPreventiveRequest{
		Search:     "%" + request.Search + "%",
		Status:     "%" + request.Status + "%",
		AssignedTo: "%" + request.AssignedTo + "%",
		StartDate:  request.StartDate,
		EndDate:    request.EndDate,
	}).Find(&list_visit_date).Error

	return list_visit_date, error
}

func (repo *repository) UpdatePreventive(request model.UpdatePreventiveRequest) (entity.Preventive, error) {
	var preventive entity.Preventive

	error := repo.db.Raw("UPDATE preventive SET visit_date = @VisitDate, location = @Location, terminal_id = @TerminalId, assigned_to = @AssignedTo, updated_by = @UpdatedBy, updated_at = @UpdatedAt, status = @Status WHERE prev_code = @PrevCode RETURNING preventive.*", entity.Preventive{
		VisitDate:  request.VisitDate,
		Location:   request.Location,
		TerminalId: request.TerminalId,
		AssignedTo: request.AssignedTo,
		UpdatedBy:  request.UpdatedBy,
		UpdatedAt:  request.UpdatedAt,
		Status:     request.Status,
		PrevCode:   request.PrevCode,
	}).Find(&preventive).Error

	return preventive, error
}

func (repo *repository) GetDetailPreventive(request string) ([]model.GetPreventiveResponse, error) {
	var preventive []model.GetPreventiveResponse

	error := repo.db.Raw("SELECT preventive.*, users.name as user_name FROM preventive LEFT OUTER JOIN users ON (preventive.assigned_to = CAST(users.id AS varchar(10))) WHERE prev_code = @PrevCode", entity.Preventive{
		PrevCode: request,
	}).Find(&preventive).Error

	return preventive, error
}

func (repo *repository) CountPreventiveByStatus(request model.CountPreventiveByStatusRequest) ([]model.CountPreventiveByStatusResponse, error) {
	var status []model.CountPreventiveByStatusResponse

	error := repo.db.Raw("SELECT status, COUNT(*) AS total FROM preventive WHERE assigned_to LIKE @AssignedTo AND CAST(visit_date AS DATE) >= @StartDate AND CAST(visit_date AS DATE) <= @EndDate GROUP BY status", model.CountPreventiveByStatusRequest{
		AssignedTo: "%" + request.AssignedTo + "%",
		StartDate:  request.StartDate,
		EndDate:    request.EndDate,
	}).Find(&status).Error

	return status, error
}
