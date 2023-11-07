package handler

import (
	"golang/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

func RegisterJobs(c *gin.Context, db *gorm.DB) (models.JobOpenings, error) {
	var newComp models.JobOpenings
	if err := c.ShouldBindJSON(&newComp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Invalid JSON data"})
		return models.JobOpenings{}, nil
	}

	validate := validator.New()
	if err := validate.Struct(newComp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid company data"})
		return models.JobOpenings{}, nil
	}

	Comp, err := CreateCompany(db, newComp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Company registration failed"})
		return models.JobOpenings{}, nil
	}

	return Comp, nil
}

func CreateCompany(db *gorm.DB, newComp models.JobOpenings) (models.JobOpenings, error) {
	comp := models.JobOpenings{
		Job_Title:   newComp.Job_Title,   // string   `json:"job_title"`
		Description: newComp.Description, // string   `json:"description"`
		Min_NP:      newComp.Min_NP,      // string   `json:"min_np"`
		Budget:      newComp.Budget,      // int8     `json:"budget"`
		MinExp:      newComp.Budget,      // int16    `json:"min_exp"
	}
	if err := db.Create(&comp).Error; err != nil {
		return models.JobOpenings{}, err
	}
	return comp, nil
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
