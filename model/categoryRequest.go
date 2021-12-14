package model

import "time"

type CreateCategoryRequest struct {
	Name             string    `json:"name"`
	CodeLevel        string    `json:"code_level"`
	Parent           string    `json:"parent"`
	UpdateAt         time.Time `json:"update_at"`
	AdditionalInput1 string    `json:"additional_input_1"`
	AdditionalInput2 string    `json:"additional_input_2"`
	AdditionalInput3 string    `json:"additional_input_3"`
}

type GetCategoryRequest struct {
	IsActive string `json:"is_active"`
}
