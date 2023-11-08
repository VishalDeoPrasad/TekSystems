package services

import (
	"golang/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JobComparison(c *gin.Context, jobs []models.JobOpenings) (models.JobOpenings, error) {
	var newComp models.JobOpenings
	if err := c.ShouldBindJSON(&newComp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Please fill up your application:"})
		return models.JobOpenings{}, nil
	}
	for _, job := range jobs {
		if job.Job_Title == newComp.Job_Title &&
			job.MinExp == newComp.MinExp && 
			job.Min_NP == newComp.Min_NP &&
			job.Budget == newComp.Budget &&
			job.Technology_Stack == newComp.Technology_Stack &&
			job.Qualification == newComp.Qualification{
			return job, nil
		}
	}
	return models.JobOpenings{}, nil
}
