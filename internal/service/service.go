package service

import (
	"context"

	"github.com/brianepv1/sv-microservice-login-golang/internal/models"
	"github.com/brianepv1/sv-microservice-login-golang/internal/repository"
)

// Service is the business logic of the application
//
//go:generate mockery --name=Service --output=service --inpackage
type Service interface {
	RegisterUser(ctx context.Context, email string, name string, password string) error
	LoginUser(ctx context.Context, email string, password string) (*models.User, error)
}

type serv struct {
	repo repository.Repository
}

func New(repo repository.Repository) Service {
	return &serv{
		repo: repo,
	}
}
