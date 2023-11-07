package models

import "gorm.io/gorm"


type JobOpenings struct {
	gorm.Model
	Job_Title        string   `json:"job_title"`
	Description      string   `json:"description"`
	Min_NP           string   `json:"min_np"`
	Budget           int16    `json:"budget"`
	MinExp           int16    `json:"min_exp"`
}

/*
type JobOpenings struct {
	gorm.Model
	Job_Title        string   `json:"job_title"`
	Min_NP           string   `json:"min_np"`
	Max_NP           string   `json:"max_np"`
	Budget           int8     `json:"budget"`
	JobLocations     []string `json:"location"`
	Technology_Stack []string `json:"technology_stack"`
	WorkMode         []string `json:"work_mode"` //[Remote,OnSite, Hybrid]
	Description      string   `json:"description"`
	MinExp           int16    `json:"min_exp"`
	MaxMax           int16    `json:"max_exp"`
	Qualification    []string `json:"qualification"`
	Shift            []string `json:"shift"`    //day, night, rotational
	JobType          []string `json:"job_type"` //[full time, part time]
}



	Min-NP
	Max-NP
	Budget
	JobLocations []
	Technology Stack[]
	WorkMode - [Remote,OnSite, Hybrid]
	Description
	MinExp
	MaxMax
	Qualification-[]
	Shift - [day, night, rotational]
	JobType - [full time, part time]

*/
