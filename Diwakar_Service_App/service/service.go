package service

import (
	"golang/models"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Conn struct {
	db *gorm.DB
}

func (c *Conn) AutoMigrate() error {
	// AutoMigrate function will ONLY create tables, missing columns and missing indexes, and WON'T change existing column's type or delete unused columns
	err := c.db.AutoMigrate(&models.User{})
	if err != nil {
		log.Error().Err(err).Msg("Failed to perform database migration")
		return err
	}
	return nil
}

func (c *Conn) CreateUser(nu models.NewUserReq) (models.User, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(nu.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Msg("hash password is not generating.")
	}

	u1 := models.User{
		Name:         nu.Name,
		Email:        nu.Email,
		PasswordHash: string(hashPass),
	}
	err = c.db.Create(&u1).Error
	if err != nil {
		log.Error().Err(err).Str("user_name", u1.Name).Msg("Failed to create user")
		return
	}
}
