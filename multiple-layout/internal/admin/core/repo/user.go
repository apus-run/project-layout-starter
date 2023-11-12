package repo

import (
	"context"

	"multiple-layout/internal/admin/core/domain/entity"
	"multiple-layout/internal/admin/core/repo/dao"
	"multiple-layout/pkg/logger"
)

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
