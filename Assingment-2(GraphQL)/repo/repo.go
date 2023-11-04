package repo

import (
	"errors"
	"golang/models"

	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

type UserRepo interface {
	CreateUser(userData models.User) (models.User, error)
	CreateCompany(companyDetails models.Company) (models.Company, error)
	ViewAllCompanies() ([]*models.Company, error)
	ViewCompanyByID(cid string)(models.Company,error)
	CreateJob(models.Job)(models.Job,error)
	ViewAllJobs()([]*models.Job,error)
	ViewJobById(id string)(models.Job,error)
	ViewJobByCid(cid string)([]*models.Job,error)
	
}

func NewRepository(db *gorm.DB) (UserRepo, error) {
	if db == nil {
		return nil, errors.New("db cannot be null")
	}
	return &Repo{
		DB: db,
	}, nil
}