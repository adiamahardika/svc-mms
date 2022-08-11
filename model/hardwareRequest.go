package model

type GetHardwareRequest struct {
	Search   string `json:"search"`
	IsActive string `json:"is_active"`
}
