package repository

import (
	"context"

	"github.com/brianepv1/sv-microservice-login-golang/internal/entity"
)

const (
	queryInsertUser = `
		INSERT INTO users 
		(	email, 
			name, 
			password, 
			lastname, 
			profile_picture, 
			address, 
			remember_token,
			created_at, 
			updated_at)
		values 
		(	?, 
			?, 
			?, 
			'testing', 
			'fake_image', 
			'Limon 36', 
			default, 
			default, 
			default);
	`

	queryGetUserByEmail = `
		SELECT id, email, name, password from users where email = ?
	`
)

func (r *repo) SaveUser(ctx context.Context, email string, name string, password string) error {
	_, err := r.db.ExecContext(ctx, queryInsertUser, email, name, password)
	return err
}

func (r *repo) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	u := &entity.User{}
	err := r.db.GetContext(ctx, u, queryGetUserByEmail, email)
	if err != nil {
		return nil, err
	}

	return u, err
}
