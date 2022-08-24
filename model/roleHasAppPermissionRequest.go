package model

type CreateRoleHasAppPermissionRequest struct {
	IdRole       int `json:"id_role"`
	IdPermission int `json:"id_permission"`
}
