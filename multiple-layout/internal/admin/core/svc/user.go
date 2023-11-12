package svc

import (
	"context"

	"multiple-layout/internal/admin/core/domain/entity"
	"multiple-layout/internal/admin/core/repo"
	"multiple-layout/pkg/logger"
)

type UserService interface {
	GetByID(ctx context.Context, id uint64) (*entity.User, error)
}

type userService struct {
	repo repo.UserRepository

	log logger.Logger
}

func NewUserService(repo repo.UserRepository, log logger.Logger) UserService {
	return &userService{
		repo: repo,
		log:  log,
	}
}

func (u userService) GetByID(ctx context.Context, id uint64) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}
