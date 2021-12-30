package model

type CreatePreventiveRequest struct {
	VisitDate  string `json:"visit_date"`
	TerminalId string `json:"terminal_id"`
	AssignedTo string `json:"assigned_to"`
	Status     string `json:"status"`
	CreatedBy  string `json:"created_by"`
}
