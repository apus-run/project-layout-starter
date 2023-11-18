package dao

import (
	"context"

	"gorm.io/gorm"

	"basic-layout/internal/core/repo/dao/model"
)

// ErrDataNotFound 通用的数据没找到
var ErrDataNotFound = gorm.ErrRecordNotFound

// UserDAO ... 数据访问层相关接口定义, 使用 DB 风格的命名
//
//go:generate mockgen -source=./user.go -package=daomocks -destination=mocks/user.mock.go UserDAO
type UserDAO interface {
	FindByID(ctx context.Context, id uint64) (*model.User, error)
}

type userDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) UserDAO {
	return &userDAO{
		db: db,
	}
}

func (u userDAO) FindByID(ctx context.Context, id uint64) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}
