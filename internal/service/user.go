package service

import (
	"context"
	"github.com/0ndreu/documents/internal/models"
)

func (s *service) CreateUser(ctx context.Context, user *models.User) error {
	panic("implement me")
}

func (s *service) UpdateUser(ctx context.Context, user *models.User) error {
	panic("implement me")
}

func (s *service) DeleteUser(ctx context.Context, ID int64) error {
	panic("implement me")
}

func (s *service) CheckUser(ctx context.Context, user *models.User) (bool, error) {
	panic("implement me")
}
