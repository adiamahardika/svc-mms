package model

type CreateRoleHasWebPermissionRequest struct {
	IdRole       int `json:"id_role"`
	IdPermission int `json:"id_permission"`
}
