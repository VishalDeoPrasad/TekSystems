package main

import (
	"context"
	"fmt"
	"golang/database"
	"golang/handler"
	"golang/middleware"
	"golang/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func startApp() (*gorm.DB, error) {
	//-------------------------Step 1: Start Database------------------------------------
	log.Info().Msg("main : Started : Initializing db support")
	db, err := database.Open()
	if err != nil {
		return nil, fmt.Errorf("connecting to db %w", err)
	}
	//pg : This variable will hold the pointer to the underlying PostgreSQL database connection.
	pg, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %w ", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	err = pg.PingContext(ctx)
	/*
		ctx: A context can be used to set a timeout for the operation. If the ping
		operation takes longer than the context's timeout, it can be canceled, and
		an error will be returned.*/
	if err != nil {
		return nil, fmt.Errorf("database is not connected: %w ", err)
	}

	//---------------------Step2:Initialize Conn layer support--------------------------
	ms, err := services.NewConn(db)
	if err != nil {
		return nil, err
	}

	//Job table create in database
	err = ms.AutoMigrate()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	db, err := startApp()
	if err != nil {
		log.Panic().Err(err).Send()
	}
	log.Info().Msg("hello, This is our Job Portal Application:")

	server := gin.New()
	server.Use(gin.Recovery(), middleware.Logger(), middleware.Auth())

	server.POST("/addJobs", func(c *gin.Context) {
		comp, err := handler.RegisterJobs(c, db)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusCreated, comp)
		}
	})

	server.GET("/view_jobs", func(c *gin.Context) {
		job_list, err := handler.ViewJobs(c, db)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Problem in fetching Data from DB"})
		} else {
			c.JSON(http.StatusOK, job_list)
		}
	})

	server.GET("/job_application_form", func(c *gin.Context) {
		c.JSON(http.StatusBadGateway, gin.H{"message": "yet to implement job_application_form"})
	})

	server.Run(":8080")
}
