package repo

import (
	"context"

	"basic-layout/internal/core/domain/entity"
	"basic-layout/internal/core/repo/dao"
	"basic-layout/pkg/logger"
)

//go:generate mockgen -source=./user.go -package=repomocks -destination=mocks/user.mock.go UserRepository
type UserRepository interface {
	FindByID(ctx context.Context, id uint64) (*entity.User, error)
}

// userRepository 使用了缓存
type userRepository struct {
	dao dao.UserDAO
	log logger.Logger
}

func NewUserRepository(dao dao.UserDAO, log logger.Logger) UserRepository {
	return &userRepository{
		dao: dao,
		log: log,
	}
}

func (u userRepository) FindByID(ctx context.Context, id uint64) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}
