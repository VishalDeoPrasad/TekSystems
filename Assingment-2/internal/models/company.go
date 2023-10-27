package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name string `json:"companyName" validate:"required"`
	City string `json:"city" validate:"required"`
}
type Job struct {
	gorm.Model
	Name       string `json:"jobTitle"`
	Field      string `json:"field"`
	Experience uint   `json:"experience"`
}
