package repository

import (
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/model"
)

type PreventiveRepositoryInterface interface {
	CreatePreventive(request entity.Preventive) (entity.Preventive, error)
	GetPreventive(request model.GetPreventiveRequest) ([]model.GetPreventiveResponse, error)
	UpdatePreventive(request model.UpdatePreventiveRequest) (entity.Preventive, error)
}

func (repo *repository) CreatePreventive(request entity.Preventive) (entity.Preventive, error) {
	var preventive entity.Preventive

	error := repo.db.Table("preventive").Create(&request).Error

	return preventive, error
}

func (repo *repository) GetPreventive(request model.GetPreventiveRequest) ([]model.GetPreventiveResponse, error) {
	var preventive []model.GetPreventiveResponse

	error := repo.db.Raw("SELECT * FROM (SELECT preventive.*, users.name AS user_name FROM preventive LEFT OUTER JOIN users ON (preventive.assigned_to = CAST(users.id AS varchar(10))) WHERE status LIKE @Status AND assigned_to LIKE @AssignedTo AND visit_date >= @StartDate AND visit_date <= @EndDate ORDER BY visit_date DESC) AS tbl WHERE terminal_id LIKE @Search", model.GetPreventiveRequest{
		Search:     "%" + request.Search + "%",
		Status:     "%" + request.Status + "%",
		AssignedTo: "%" + request.AssignedTo + "%",
		StartDate:  request.StartDate,
		EndDate:    request.EndDate,
	}).Find(&preventive).Error

	return preventive, error
}

func (repo *repository) UpdatePreventive(request model.UpdatePreventiveRequest) (entity.Preventive, error) {
	var preventive entity.Preventive

	error := repo.db.Raw("UPDATE preventive SET visit_date = @VisitDate, terminal_id = @TerminalId, assigned_to = @AssignedTo, updated_by = @UpdatedBy, updated_at = @UpdatedAt, status = @Status WHERE prev_code = @PrevCode RETURNING preventive.*", entity.Preventive{
		VisitDate:  request.VisitDate,
		TerminalId: request.TerminalId,
		AssignedTo: request.AssignedTo,
		UpdatedBy:  request.UpdatedBy,
		UpdatedAt:  request.UpdatedAt,
		Status:     request.Status,
		PrevCode:   request.PrevCode,
	}).Find(&preventive).Error

	return preventive, error
}
