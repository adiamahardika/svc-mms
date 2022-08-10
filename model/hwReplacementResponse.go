package model

import "svc-monitoring-maintenance/entity"

type GetHwReplacementResponse struct {
	TicketCode    string                 `json:"ticket_code"`
	NoSPM         string                 `json:"no_spm"`
	NoReqSPM      string                 `json:"no_req_spm"`
	HwReplacement []entity.HwReplacement `json:"hw_replacement"`
}
