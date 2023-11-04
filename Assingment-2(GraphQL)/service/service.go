package service

import (
	"errors"
	"golang/graph/model"
	"golang/repo"
)

type UserService interface {
	UserSignup(model.NewUser) (*model.User, error)
	CreateCompany(model.NewCompany) (*model.Company, error)
	ViewAllCompanies() ([]*model.Company, error)
	ViewCompanyByID(cid string) (*model.Company, error)
	CreateJob(model.NewJob) (*model.Job, error)
	ViewAllJobs() ([]*model.Job, error)
	ViewJobByID(id string) (*model.Job, error)
	ViewJobByCid(cid string)([]*model.Job,error)
}

type Service struct {
	userRepo repo.UserRepo
}

func NewService(userRepo repo.UserRepo) (UserService, error) {
	if userRepo == nil {
		return nil, errors.New("interface cannot be null")
	}
	return &Service{
		userRepo: userRepo,
	}, nil
}
