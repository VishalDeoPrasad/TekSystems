package services

import (
	"context"
	"golang/internal/models"

	"github.com/golang-jwt/jwt/v5"
)

//go:generate mockgen -source service.go -destination mockmodels/service_mock.go -package mockmodels

type Service interface {
	// CreatInventory(ctx context.Context, ni NewInventory, userId uint) (Inventory, error)
	// ViewInventory(ctx context.Context, userId string) ([]Inventory, float64, error)
	CreateCompany(ctx context.Context, newComp models.Company) (models.Company, error) 
	CreateUser(ctx context.Context, nu models.NewUser) (models.User, error)
	Authenticate(ctx context.Context, email, password string) (jwt.RegisteredClaims,
		error)
	AutoMigrate() error
}

type Store struct {
	Service
}

func NewStore(s Service) Store {
	return Store{Service: s}
}
