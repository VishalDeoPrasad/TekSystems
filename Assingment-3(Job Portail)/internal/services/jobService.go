// // * Job Service
package services

import (
	"context"
	"golang/internal/models"
	"strconv"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func intToLocation(slice []uint) []models.Location {
	var locRespo []models.Location
	for _, s := range slice {
		loc := models.Location{
			Model: gorm.Model{ID: s},
		}
		locRespo = append(locRespo, loc)
	}
	return locRespo
}

func intToSkill(slice []uint) []models.Skill {
	var locRespo []models.Skill
	for _, s := range slice {
		loc := models.Skill{
			Model: gorm.Model{ID: s},
		}
		locRespo = append(locRespo, loc)
	}
	return locRespo
}

func intToQualification(slice []uint) []models.Qualification {
	var locRespo []models.Qualification
	for _, s := range slice {
		loc := models.Qualification{
			Model: gorm.Model{ID: s},
		}
		locRespo = append(locRespo, loc)
	}
	return locRespo
}

/*func (s *DbConnStruct) JobByCompanyId(jobs []models.JobReq, compId string) ([]models.JobRespo, error) {
	companyId, err := strconv.ParseUint(compId, 10, 64)
	if err != nil {
		return nil, err
	}
	var job1 []models.JobRespo

	for _, j := range jobs {
		job := models.Job{
			Name:           j.Name,
			Field:          j.Field,
			Experience:     j.Experience,
			Min_NP:         j.Min_NP,
			Max_NP:         j.Max_NP,
			Budget:         j.Budget,
			Locations:      intToLocation(j.Locations),
			Skill:          intToSkill(j.Skill),
			WorkMode:       j.WorkMode,
			Description:    j.Description,
			MinExp:         j.MinExp,
			Qualifications: intToQualification(j.Qualifications),
			Shift:          j.Shift,
			CompanyId:      companyId,
		}
		err := s.db.Create(&job).Error
		if err != nil {
			return nil, err
		}
		jData := models.JobRespo{
			Id: job.ID,
		}
		job1 = append(job1, jData)
	}

	return job1, nil
}*/

func (s *DbConnStruct) JobByCompanyId(jobR models.JobReq, compId string) (models.JobRespo, error) {
	companyId, err := strconv.ParseUint(compId, 10, 64)
	if err != nil {
		log.Info().Msg("main : Started : Problem with parseUint(compId)")
		return models.JobRespo{}, err
	}
	job := models.Job{
		Name:           jobR.Name,
		Field:          jobR.Field,
		Experience:     jobR.Experience,
		Min_NP:         jobR.Min_NP,
		Max_NP:         jobR.Max_NP,
		Budget:         jobR.Budget,
		Locations:      intToLocation(jobR.Locations),
		Skill:          intToSkill(jobR.Skill),
		WorkMode:       jobR.WorkMode,
		Description:    jobR.Description,
		MinExp:         jobR.MinExp,
		Qualifications: intToQualification(jobR.Qualifications),
		Shift:          jobR.Shift,
		CompanyId:      companyId,
	}
	err = s.db.Create(jobR).Error
	if err != nil {
		log.Info().Msg("main : Started : Problem in Create(&job)")
		return models.JobRespo{}, err
	}
	jData := models.JobRespo{
		Id: job.ID,
	}

	return jData, nil
}

func (s *DbConnStruct) FetchJobByCompanyId(ctx context.Context, companyId string) ([]models.Job, error) {
	var listOfJobs []models.Job
	tx := s.db.WithContext(ctx).Where("company_id = ?", companyId)
	err := tx.Find(&listOfJobs).Error
	if err != nil {
		return nil, err
	}

	return listOfJobs, nil
}

func (s *DbConnStruct) GetJobById(ctx context.Context, jobId string) (models.Job, error) {
	var jobData models.Job
	tx := s.db.WithContext(ctx).Where("ID = ?", jobId)
	err := tx.Find(&jobData).Error
	if err != nil {
		return models.Job{}, err
	}

	return jobData, nil
}

func (s *DbConnStruct) GetAllJobs(ctx context.Context) ([]models.Job, error) {
	var listJobs []models.Job
	tx := s.db.WithContext(ctx)
	err := tx.Find(&listJobs).Error
	if err != nil {
		return nil, err
	}

	return listJobs, nil
}
