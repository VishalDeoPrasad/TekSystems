package handlers

import (
	"fmt"
	"golang/internal/auth"
	"golang/internal/middleware"
	"golang/internal/services"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
)

func API(a *auth.Auth, c *services.DbConnStruct) *gin.Engine {

	// Create a new Gin engine; Gin is a HTTP web framework written in Go
	ginEngine := gin.New()

	// Attempt to create new middleware with authentication
	// Here, *auth.Auth passed as a parameter will be used to set up the middleware
	mid, err := middleware.NewMid(a)
	ms := services.NewStore(c)
	h := handler{
		s: ms,
		a: a,
	}

	// If there is an error in setting up the middleware, panic and stop the application
	// then log the error message
	if err != nil {
		log.Panic().Msg("middlewares not set up")
	}

	// Attach middleware's Log function and Gin's Recovery middleware to our application
	// The Recovery middleware recovers from any panics and writes a 500 HTTP response if there was one.
	ginEngine.Use(middleware.Logger(), gin.Recovery())

	// Define a route at path "/check"
	// If it receives a GET request, it will use the m.Authenticate(check) function.
	ginEngine.GET("/check_token", mid.Authenticate(check))
	ginEngine.POST("/signup_page", h.Signup)
	ginEngine.POST("/login_page", h.Login)
	ginEngine.POST("/register_company", h.RegisterCompany)
	ginEngine.GET("/list_all_companies", h.fetchListOfCompany)
	ginEngine.GET("/fatch_company_by_id/:ID", h.companyById)
	ginEngine.POST("/addJobs/:ID", h.addJobsById)
	ginEngine.GET("/fetch_jobs_by_compID/:companyId", h.jobsByCompanyById)
	ginEngine.GET("/fetchAllJob", h.GetAllJobs)
	ginEngine.GET("/fetchJob/:ID", h.fetchJobById)

	// Return the prepared Gin engine
	return ginEngine
}

func check(c *gin.Context) {
	select {
	case <-c.Request.Context().Done():
		fmt.Println("user not there")
		return
	default:
		c.JSON(http.StatusOK, gin.H{"msg": "statusOk"})
	}
}
