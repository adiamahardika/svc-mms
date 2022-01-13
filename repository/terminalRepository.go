package repository

import (
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/model"
)

type TerminalRepositoryInterface interface {
	GetTerminal(request model.GetTerminalRequest) ([]entity.Terminal, error)
}

func (repo *repository) GetTerminal(request model.GetTerminalRequest) ([]entity.Terminal, error) {
	var terminal []entity.Terminal

	error := repo.db.Raw("SELECT * FROM (SELECT ms_terminal.* FROM ms_terminal WHERE area LIKE @Area AND regional LIKE @Regional AND ctp_type LIKE @CtpType AND kecamatan LIKE @Kecamatan AND kota_kabupaten LIKE @KotaKabupaten ORDER BY terminal_id ASC) AS tbl WHERE tbl.terminal_id LIKE @Search OR tbl.terminal_name LIKE @Search OR tbl.terminal_location LIKE @Search OR tbl.pic LIKE @Search", model.GetTerminalRequest{
		Search:        "%" + request.Search + "%",
		Area:          "%" + request.Area + "%",
		Regional:      "%" + request.Regional + "%",
		CtpType:       "%" + request.CtpType + "%",
		Kecamatan:     "%" + request.Kecamatan + "%",
		KotaKabupaten: "%" + request.KotaKabupaten + "%",
	}).Find(&terminal).Error

	return terminal, error
}
