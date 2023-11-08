package services

import (
	"golang/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JobComparison(c *gin.Context, jobs []models.Job_Portal) (models.Job_Portal, error) {
	var newComp models.Job_Portal
	if err := c.ShouldBindJSON(&newComp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Please fill up your application:"})
		return models.Job_Portal{}, nil
	}
	for _, job := range jobs {
		if job.Name == newComp.Name && 
			job.Budget == newComp.Budget && 
			job.Min_NP == newComp.Min_NP && 
			job.Max_NP == newComp.Max_NP {
			return job, nil
		}
	}
	return models.Job_Portal{}, nil
}
