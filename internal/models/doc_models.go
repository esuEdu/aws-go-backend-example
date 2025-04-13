package models

import "gorm.io/gorm"

type Doc struct {
	gorm.Model
	Name string `json:"name"`
	Url  string `json:"url"`
}
