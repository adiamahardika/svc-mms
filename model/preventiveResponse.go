package model

import "svc-monitoring-maintenance/entity"

type GetVisitDateResponse struct {
	VisitDate       string `json:"visit_date"`
	TotalPreventive string `json:"total_preventive"`
}

type GetGroupPreventiveResponse struct {
	VisitDate       string              `json:"visit_date"`
	TotalPreventive string              `json:"total_preventive"`
	PreventiveList  []entity.Preventive `json:"preventive_list"`
}

type CountPreventiveByStatusResponse struct {
	Status string `json:"status"`
	Total  int    `json:"total"`
}
