package service

import (
	"context"
	"errors"

	encryption "github.com/brianepv1/sv-microservice-login-golang/encryption"
	"github.com/brianepv1/sv-microservice-login-golang/internal/models"
)

var (
	ErrUserAlreadyExists = errors.New("user already exits")
	PwsUserIsWrong       = errors.New("password is wrong")
)

func (s *serv) RegisterUser(ctx context.Context, email string, name string, password string) error {
	u, _ := s.repo.GetUserByEmail(ctx, email)
	if u != nil {
		return ErrUserAlreadyExists
	}

	bb, err := encryption.Encrypt([]byte(password))
	if err != nil {
		return err
	}

	pass := encryption.ToBase64(bb)

	return s.repo.SaveUser(ctx, email, name, pass)
}

func (s *serv) LoginUser(ctx context.Context, email string, password string) (*models.User, error) {
	u, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	bb, err := encryption.FromBase64(u.Password)
	if err != nil {
		return nil, err
	}

	decryptedPassword, err := encryption.Decrypt(bb)
	if err != nil {
		return nil, err
	}

	if string(decryptedPassword) != password {
		return nil, PwsUserIsWrong
	}

	return &models.User{
		ID:    u.ID,
		Email: u.Email,
		Name:  u.Name,
	}, err
}
