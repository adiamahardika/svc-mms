package model

import (
	"svc-monitoring-maintenance/entity"
)

type ChecklistPreventiveResponse struct {
	Header *entity.HeaderChecklistPreventive `json:"header" form:"header"`
	Items  []*entity.ChecklistHw             `json:"items" form:"items"`
	User   []*entity.UserChecklistPreventive `json:"user" form:"user"`
}
