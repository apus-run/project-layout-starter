package repo

import (
	"context"

	"github.com/apus-run/sea-kit/log"

	"gaia-layout/internal/core/domain/entity"
	"gaia-layout/internal/core/repo/dao"
)

type UserRepository interface {
	FindByID(ctx context.Context, id uint64) (*entity.User, error)
}

// userRepository 使用了缓存
type userRepository struct {
	dao dao.UserDAO
	log log.Logger
}

func NewUserRepository(dao dao.UserDAO, log log.Logger) UserRepository {
	return &userRepository{
		dao: dao,
		log: log,
	}
}

func (u userRepository) FindByID(ctx context.Context, id uint64) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}
