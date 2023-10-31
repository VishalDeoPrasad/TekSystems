package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name string `json:"name" validate:"required"`
	City string `json:"city" validate:"required"`
	Jobs []Job  `json:"jobs,omitempty" gorm:"foreignKey:CompanyId"`
}
type Job struct {
	gorm.Model
	Name       string `json:"title"`
	Field      string `json:"field"`
	Experience uint   `json:"experience"`
	CompanyId  uint64 `json:"companyId"`
}
