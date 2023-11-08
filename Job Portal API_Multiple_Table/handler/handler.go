package handler

import (
	"golang/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterJobs(c *gin.Context, db *gorm.DB) (models.JobOpenings, error) {
	var newComp models.JobOpenings
	c.BindJSON(&newComp)

	// validate := validator.New()
	// if err := validate.Struct(newComp); err != nil {
	// 	log.Error().Err(err).Msg("Problem in validation")
	// 	c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid company data"})
	// 	return models.JobOpenings{}, err
	// }

	Comp, err := CreateCompany(db, newComp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Company registration failed"})
		return models.JobOpenings{}, err
	}

	return Comp, nil
}

func CreateCompany(db *gorm.DB, newJob models.JobOpenings) (models.JobOpenings, error) {
	// comp := models.JobOpenings{
	// 	// Job_Title:   newComp.Job_Title,   // string   `json:"job_title"`
	// 	// Description: newComp.Description, // string   `json:"description"`
	// 	// Min_NP:      newComp.Min_NP,      // string   `json:"min_np"`
	// 	// Budget:      newComp.Budget,      // int8     `json:"budget"`

	// 	Job_Title:        newJob.Job_Title,
	// 	Min_NP:           newJob.Min_NP,
	// 	Max_NP:           newJob.Max_NP,
	// 	Budget:           newJob.Budget,
	// 	Description:      newJob.Description,
	// 	MinExp:           newJob.MinExp,
	// 	MaxMax:           newJob.MaxMax,
	// 	JobLocations:     newJob.JobLocations,
	// 	Technology_Stack: newJob.Technology_Stack,
	// 	WorkMode:         newJob.WorkMode,
	// 	Qualification:    newJob.Qualification,
	// 	Shift:            newJob.Shift,
	// 	JobType:          newJob.JobType,
	// }
	if err := db.Create(&newJob).Error; err != nil {
		return models.JobOpenings{}, err
	}
	return newJob, nil
}

func ViewJobs(c *gin.Context, db *gorm.DB) ([]models.JobOpenings, error) {
	var listComp []models.JobOpenings
	tx := db.WithContext(c)
	if err := tx.Find(&listComp).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return []models.JobOpenings{}, nil
	}
	return listComp, nil
}
