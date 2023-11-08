package handler

import (
	"golang/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

func RegisterJobs(c *gin.Context, db *gorm.DB) (models.Job_Portal, error) {
	var newComp models.Job_Portal
	c.BindJSON(&newComp)

	validate := validator.New()
	if err := validate.Struct(newComp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid Job data"})
		return models.Job_Portal{}, err
	}

	Comp, err := CreateJob(db, newComp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Job registration failed"})
		return models.Job_Portal{}, err
	}

	return Comp, nil
}

func CreateJob(db *gorm.DB, newJob models.Job_Portal) (models.Job_Portal, error) {
	// comp := models.Job{
	// 	JobTitle:        newJob.JobTitle,
	// 	JobSalary:       newJob.JobSalary,
	// 	Company:         newJob.Company,
	// 	Uid:             newJob.Uid,
	// 	MinNoticePeriod: newJob.MinNoticePeriod,
	// 	MaxNoticePeriod: newJob.MaxNoticePeriod,
	// 	Budget:          newJob.Budget,
	// 	Description:     newJob.Description,
	// 	MinExperience:   newJob.MinExperience,
	// 	MaxExperience:   newJob.MaxExperience,
	// 	JobLocations:    newJob.JobLocations,
	// 	Technology:      newJob.Technology,
	// 	WorkMode:        newJob.WorkMode,
	// 	Qualification:   newJob.Qualification,
	// 	Shift:           newJob.Shift,
	// 	JobType:         newJob.JobType,
	// }

	if err := db.Create(&newJob).Error; err != nil {
		return models.Job_Portal{}, err
	}
	return newJob, nil
}

func ViewJobs(c *gin.Context, db *gorm.DB) ([]models.Job_Portal, error) {
	var listComp []models.Job_Portal
	tx := db.WithContext(c)
	if err := tx.Find(&listComp).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return []models.Job_Portal{}, nil
	}
	return listComp, nil
}
