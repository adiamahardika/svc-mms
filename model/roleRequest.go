package model

type CreateRoleRequest struct {
	Name string `json:"name"`
}

type GetRoleRequest struct {
	IsActive string `json:"is_active"`
}
