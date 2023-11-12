package svc

import (
	"context"
	"github.com/apus-run/sea-kit/log"

	"gaia-layout/internal/core/domain/entity"
	"gaia-layout/internal/core/repo"
)

type UserService interface {
	GetByID(ctx context.Context, id uint64) (*entity.User, error)
	UpdateUser(ctx context.Context, userEntity entity.User) (*entity.User, error)
}

type userService struct {
	repo repo.UserRepository

	log log.Logger
}

func NewUserService(repo repo.UserRepository, log log.Logger) UserService {
	return &userService{
		repo: repo,
		log:  log,
	}
}

func (u *userService) GetByID(ctx context.Context, id uint64) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userService) UpdateUser(ctx context.Context, userEntity entity.User) (*entity.User, error) {
	// 从 FromContext(ctx) 中拿到 ID
	// 根据 ID 查询用户
	// 更新用户

	return &entity.User{
		ID:    0,
		Name:  "moocss",
		Email: "moocss@163.com",
	}, nil
}
