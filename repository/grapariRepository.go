package repository

import (
	"fmt"
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/model"
)

type GrapariRepositoryInterface interface {
	GetGrapari(request model.GetGrapariRequest) ([]entity.MsGrapari, error)
}

func (repo *repository) GetGrapari(request model.GetGrapariRequest) ([]entity.MsGrapari, error) {
	var result []entity.MsGrapari
	var area_code string
	var regional string
	var grapari_id string

	if len(request.AreaCode) > 0 {
		area_code = "area IN @AreaCode AND"
	}
	if len(request.Regional) > 0 {
		regional = "regional IN @Regional AND"
	}
	if len(request.GrapariId) > 0 {
		grapari_id = "grapari_id IN @GrapariId AND "
	}

	query := fmt.Sprintf("SELECT * FROM ms_grapari WHERE %s %s %s status LIKE @Status AND name LIKE @Search ORDER BY name ASC", area_code, regional, grapari_id)

	error := repo.db.Raw(query, model.GetGrapariRequest{
		AreaCode:  request.AreaCode,
		Regional:  request.Regional,
		GrapariId: request.GrapariId,
		Search:    "%" + request.Search + "%",
		Status:    "%" + request.Status + "%",
	}).Find(&result).Error

	return result, error
}
