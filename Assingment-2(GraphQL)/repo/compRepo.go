package repo

import (
	"errors"
	"golang/models"

	"github.com/rs/zerolog/log"
)

func (r *Repo) CreateCompany(companyDetails models.Company) (models.Company, error) {
	result := r.DB.Create(&companyDetails)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return models.Company{}, errors.New("could not create the record in the table")
	}
	return companyDetails, nil
}

func (r *Repo) ViewAllCompanies() ([]*models.Company, error) {
	var companyDetails []*models.Company
	result := r.DB.Find(&companyDetails)
	if result.Error != nil {
		return nil, errors.New("could not fetch the data")
	}
	return companyDetails, nil
}

func (r *Repo) ViewCompanyByID(cid string) (models.Company, error) {
	var companyDetails models.Company
	result := r.DB.Where("id = ?", cid).First(&companyDetails)
	if result.Error != nil {
		return models.Company{}, errors.New("could not find the company in the records")
	}
	return companyDetails, nil
}

func (r *Repo) CreateJob(jd models.Job) (models.Job, error) {
	result := r.DB.Create(&jd)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return models.Job{}, errors.New("could not create the record in the table")
	}
	return jd, nil
}
func (r *Repo) ViewAllJobs() ([]*models.Job, error) {
	var jobDetails []*models.Job
	result := r.DB.Find(&jobDetails)
	if result.Error != nil {
		return nil, errors.New("could not fetch the data")
	}
	return jobDetails, nil
}

func (r *Repo) ViewJobById(id string) (models.Job, error) {
	var jobData models.Job
	result := r.DB.Where("id = ?", id).First(&jobData)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return models.Job{}, errors.New("could not find the company")
	}
	return jobData, nil

}
func (r *Repo) ViewJobByCid(cid string) ([]*models.Job, error) {
	var jobData []*models.Job
	result := r.DB.Where("id = ?", cid).First(&jobData)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return []*models.Job{}, errors.New("could not find the company")
	}
	return jobData, nil

}
