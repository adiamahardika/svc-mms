package repository

import (
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/model"
)

type GrapariRepositoryInterface interface {
	GetGrapari(request model.GetGrapariRequest) ([]entity.Grapari, error)
}

func (repo *repository) GetGrapari(request model.GetGrapariRequest) ([]entity.Grapari, error) {
	var grapari []entity.Grapari

	error := repo.db.Raw("SELECT * FROM (SELECT ms_grapari.* FROM ms_grapari WHERE area LIKE @Area AND regional LIKE @Regional AND status LIKE @Status ORDER BY name ASC) AS tbl WHERE tbl.name LIKE @Search OR tbl.grapari_id LIKE @Search", model.GetGrapariRequest{
		Search:   "%" + request.Search + "%",
		Area:     "%" + request.Area + "%",
		Regional: "%" + request.Regional + "%",
		Status:   "%" + request.Status + "%",
	}).Find(&grapari).Error

	return grapari, error
}
