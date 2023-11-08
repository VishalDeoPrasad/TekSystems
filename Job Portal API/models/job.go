package models

import "gorm.io/gorm"

/*
type JobOpenings struct {
	gorm.Model
	Job_Title        string   `json:"job_title"`
	Min_NP           string   `json:"min_np"`
	Max_NP           string   `json:"max_np"`
	Budget           int8     `json:"budget"`
	JobLocations     string `json:"location"`
	Technology_Stack string `json:"technology_stack"`
	WorkMode         string `json:"work_mode"` //[Remote,OnSite, Hybrid]
	Description      string   `json:"description"`
	MinExp           int16    `json:"min_exp"`
	MaxMax           int16    `json:"max_exp"`
	Qualification    string `json:"qualification"`
	Shift            string `json:"shift"`    //day, night, rotational
	JobType          string `json:"job_type"` //[full time, part time]
}
*/

/*
type JobOpenings struct {
	gorm.Model
	Job_Title   string `json:"job_title"`
	Description string `json:"description"`
	Min_NP      string `json:"min_np"`
	Budget      int16  `json:"budget"`
	MinExp      int16  `json:"min_exp"`
}
*/

type Job struct {
	gorm.Model
	JobTitle        string          `json:"job_title" validate:"required"`
	JobSalary       string          `json:"job_salary" validate:"required"`
	Company         Company         `gorm:"ForeignKey:uid"`
	Uid             uint64          `JSON:"uid, omitempty"`
	MinNoticePeriod string          `json:"min_np"  validate:"required"`
	MaxNoticePeriod string          `json:"max_np"  validate:"required"`
	Budget          string          `json:"budget" validate:"required"`
	Description     string          `json:"description" validate:"required"`
	MinExperience   string          `json:"Min_exp" validate:"required"`
	MaxExperience   string          `json:"Max_exp" validate:"required"`
	Locations       []Location      `gorm:"many2many:job_locations" json:"job_locations"`
	Technologies    []Technology    `gorm:"many2many:job_techs" json:"technologies"`
	WorkModes       []WorkMode      `gorm:"many2many:job_workmode" json:"work_mode" validate:"required"`
	Qualifications  []Qualification `gorm:"many2many:job_qualifications" json:"qualification"`
	Shifts          []Shift         `gorm:"many2many:job_shifts" json:"job_shifts"`
	JobTypes        []JobType       `gorm:"many2many:job_jobtypes" json:"job_type"`
}

type WorkMode struct{
	gorm.Model
	Name string `json:"name" gorm:"unique"`
}
type Location struct {
	gorm.Model
	Name string `json:"name" gorm:"unique"`
}

type Technology struct {
	gorm.Model
	Name string `json:"name" gorm:"unique"`
}

type Qualification struct {
	gorm.Model
	Name string `json:"name" gorm:"unique"`
}

type Shift struct {
	gorm.Model
	Name string `json:"name" gorm:"unique"`
}

type JobType struct {
	gorm.Model
	Name string `json:"name" gorm:"unique"`
}

type Company struct {
	gorm.Model
	CompanyName string `json:"company_name"  gorm:"unique"`
	Adress      string `json:"company_adress" validate:"required"`
	Domain      string `json:"domain" validate:"required"`
}

type CreateCompany struct {
	CompanyName string `json:"company_name" validate:"required"`
	Adress      string `json:"company_adress" validate:"required"`
	Domain      string `json:"domain" validate:"required"`
}

type NewJobRequest struct {
	Name       string `json:"title" validate:"required"`
	Field      string `json:"field" validate:"required"`
	Experience uint   `json:"experience" validate:"required"`
	Min_NP     uint   `json:"min-NP" validate:"required"`
	Max_NP     uint   `json:"max-NP" validate:"required"`
	Budget     uint   `json:"salary" validate:"required"`
	Locations  []uint `json:"locations" validate:"required"`
	//Stack          []Skill         `json:"skills" validate:"required" gorm:"many2many:job_skill;"`
	WorkMode    string `json:"workMode" validate:"required"`
	Description string `json:"desc" validate:"required"`
	MinExp      uint   `json:"minExp" validate:"required"`
	//Qualifications []Qualification `json:"qualification" validate:"required" gorm:"many2many:job_qualifications;"`
	Shift     string `json:"shift" validate:"required"`
	CompanyId uint64 `json:"companyId"`
}

type NewJobResponse struct {
	ID uint
}
