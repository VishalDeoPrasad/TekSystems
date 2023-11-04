package repo

import (
	"errors"
	"golang/models"

	"github.com/rs/zerolog/log"
)

func (r *Repo) CreateUser(userDetails models.User) (models.User, error) {
	result := r.DB.Create(&userDetails)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return models.User{}, errors.New("could not create the record in the table")
	}
	return userDetails, nil
}