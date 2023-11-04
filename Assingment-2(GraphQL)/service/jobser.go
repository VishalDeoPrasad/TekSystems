package service

import (
	"strconv"

	"golang/graph/model"
	newModel "golang/models"
	"golang/password"
)

func (s *Service) UserSignup(userData model.NewUser) (*model.User, error) {
	hashedPassword, err := password.HashPassword(userData.Password)
	if err != nil {
		return nil, err
	}

	userDetails := newModel.User{
		Username:     userData.Username,
		Email:        userData.Email,
		HashPassword: hashedPassword,
	}

	userDetails, err = s.userRepo.CreateUser(userDetails)
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:        strconv.FormatUint(uint64(userDetails.ID), 10),
		Username:  userDetails.Username,
		Email:     userDetails.Email,
		CreatedAt: userDetails.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: userDetails.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}
