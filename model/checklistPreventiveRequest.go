package model

import "mime/multipart"

type CreateChecklistPreventiveRequest struct {
	Header               string                `json:"header" form:"header" binding:"required"`
	Items                string                `json:"items" form:"items" binding:"required"`
	UserTrilogi          string                `json:"user_trilogi" form:"user_trilogi" binding:"required"`
	UserTrilogiSignature *multipart.FileHeader `form:"user_trilogi_signature" binding:"required"`
	UserTsel             string                `json:"user_tsel" form:"user_tsel" binding:"required"`
	UserTselSignature    *multipart.FileHeader `form:"user_tsel_signature" binding:"required"`
}
