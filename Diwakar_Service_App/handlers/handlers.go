package handlers

import (
	"encoding/json"
	"golang/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func Signup(c *gin.Context) {
	var nu models.NewUserReq
		body := c.Request.Body
		err := json.NewDecoder(body).Decode(&nu)
		if err != nil {
			log.Error().Err(err).Msg("Problem in reading request body")
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "Invalid JSON format"})
		}

		validate := validator.New()
		err = validate.Struct(nu)
		if err != nil {
			log.Error().Err(err).Msg("please provide name, email and password.")
			c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"msg": "Validation failed"})
		}

		// hashPass, err := bcrypt.GenerateFromPassword([]byte(nu.Password), bcrypt.DefaultCost)
		// if err != nil {
		// 	log.Error().Err(err).Msg("hash password is not generating.")
		// }

		// u1 := models.User{
		// 	Name:         nu.Name,
		// 	Email:        nu.Email,
		// 	PasswordHash: string(hashPass),
		// }
		// db.AutoMigrate(&models.User{})
		// err = db.Create(&u1).Error
		// if err != nil {
		// 	log.Error().Err(err).Str("user_name", u1.Name).Msg("Failed to create user")
		// 	return
		// }
		log.Info().Str("user_name", u1.Name).Msg("User created successfully")
		c.JSON(http.StatusOK, u1)
	
}

func Login(c *gin.Context) {
	// Dummy implementation for the Login handler
	c.JSON(200, gin.H{"message": "Login Successful"})
}

func API(db *gorm.DB) *gin.Engine {
	// Create a new Gin engine
	r := gin.New()

	// Use the middleware and recovery globally
	r.Use(gin.Logger(), gin.Recovery())

	// Define routes
	r.POST("/signup", Signup)
		

	r.POST("/login", Login)

	return r
}
