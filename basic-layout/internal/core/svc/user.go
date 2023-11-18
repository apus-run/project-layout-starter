package svc

import (
	"context"

	"basic-layout/internal/core/domain/entity"
	"basic-layout/internal/core/repo"
	"basic-layout/pkg/logger"
)

//go:generate mockgen -source=./user.go -package=svcmocks -destination=mocks/user.mock.go UserService
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
