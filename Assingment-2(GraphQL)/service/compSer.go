package service

import (
	"strconv"

	"golang/graph/model"
	newModel "golang/models"
)

func (s *Service) CreateCompany(companyData model.NewCompany) (*model.Company, error) {
	companyDetails := newModel.Company{
		Name:     companyData.Name,
		Location: companyData.Location,
		Salary:   companyData.Salary,
	}

	companyDetails, err := s.userRepo.CreateCompany(companyDetails)
	if err != nil {
		return nil, err
	}

	return &model.Company{
		ID:       strconv.FormatUint(uint64(companyDetails.ID), 10),
		Name:     companyData.Name,
		Location: companyData.Location,
		Salary:   companyData.Salary,
	}, nil

}

func (s *Service) ViewAllCompanies() ([]*model.Company, error) {
	companies, err := s.userRepo.ViewAllCompanies()
	if err != nil {
		return nil, err
	}
	cpDatas := []*model.Company{}

	for _, v := range companies {
		cpData := model.Company{
			ID:       strconv.FormatUint(uint64(v.ID), 10),
			Name:     v.Name,
			Location: v.Location,
			Salary:   v.Salary,
		}
		cpDatas = append(cpDatas, &cpData)
	}
	return cpDatas, nil
}

func (s *Service) ViewCompanyByID(cid string) (*model.Company, error) {
	userDetails, err := s.userRepo.ViewCompanyByID(cid)
	if err != nil {
		return nil, err
	}

	return &model.Company{
		ID:       strconv.FormatUint(uint64(userDetails.ID), 10),
		Name:     userDetails.Name,
		Location: userDetails.Location,
		Salary:   userDetails.Salary,
	}, nil

}

func (s *Service) CreateJob(jobData model.NewJob) (*model.Job, error) {
	jobDetails := newModel.Job{
		Jobtitle:   jobData.Jobtitle,
		Experience: jobData.Experience,
		Salary:     jobData.Salary,
	}

	jobDetails, err := s.userRepo.CreateJob(jobDetails)
	if err != nil {
		return nil, err
	}

	return &model.Job{
		ID:         strconv.FormatUint(uint64(jobDetails.ID), 10),
		Jobtitle:   jobData.Jobtitle,
		Experience: jobData.Experience,
		Salary:     jobData.Salary,
	}, nil

}
func (s *Service) ViewAllJobs() ([]*model.Job, error) {
	jobs, err := s.userRepo.ViewAllJobs()
	if err != nil {
		return nil, err
	}
	jDatas := []*model.Job{}

	for _, v := range jobs {
		jData := model.Job{
			ID:         strconv.FormatUint(uint64(v.ID), 10),
			Jobtitle:   v.Jobtitle,
			Salary:     v.Salary,
			Experience: v.Experience,
		}
		jDatas = append(jDatas, &jData)
	}
	return jDatas, nil
}

func (s *Service) ViewJobByID(id string) (*model.Job, error) {
	jobData, err := s.userRepo.ViewJobById(id)
	if err != nil {
		return &model.Job{}, err
	}
	return &model.Job{
		ID:         strconv.FormatUint(uint64(jobData.ID), 10),
		Jobtitle:   jobData.Jobtitle,
		Salary:     jobData.Salary,
		Experience: jobData.Experience,
	}, nil

}
func (s *Service) ViewJobByCid(cid string) ([]*model.Job, error) {
	jobDetails, err := s.userRepo.ViewJobByCid(cid)
	if err != nil {
		return nil, err
	}
	var jobDatas []*model.Job
	for _, v := range jobDetails {
		jobData := &model.Job{
			ID:         strconv.FormatUint(uint64(v.ID), 10),
			Jobtitle:   v.Jobtitle,
			Salary:     v.Salary,
			Experience: v.Experience,
		}
		jobDatas = append(jobDatas, jobData)
	}
	return jobDatas, nil

}
