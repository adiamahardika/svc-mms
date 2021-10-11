package entity

import "time"

type Category struct {
	Id        int    `json:"id" gorm:"primaryKey"`
	CodeLevel string `json:"code_level"`
	Name      string `json:"name"`
	Parent    string `json:"parent"`
	AdditionalInput1 string `json:"additional_input_1"`
	AdditionalInput2 string `json:"additional_input_2"`
	AdditionalInput3 string `json:"additional_input_3"`
	UpdateAt  time.Time `json:"update_at"`
}