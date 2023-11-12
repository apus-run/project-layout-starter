package rpc

import (
	"context"

	v1 "gaia-layout/api/user/v1"
	"gaia-layout/internal/core/domain/entity"
	"gaia-layout/internal/core/svc"
)

var _ v1.UserServer = (*UserService)(nil)

type UserService struct {
	v1.UnimplementedUserServer

	userSvc svc.UserService
}

func NewUserService(userSvc svc.UserService) *UserService {
	return &UserService{
		userSvc: userSvc,
	}
}

func (s *UserService) Login(ctx context.Context, req *v1.LoginRequest) (*v1.UserReply, error) {
	return &v1.UserReply{}, nil
}
func (s *UserService) Register(ctx context.Context, req *v1.RegisterRequest) (*v1.UserReply, error) {
	return &v1.UserReply{}, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (*v1.UserReply, error) {
	svc, err := s.userSvc.UpdateUser(ctx, entity.User{
		Name:  req.User.Name,
		Email: req.User.Email,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UserReply{
		User: &v1.UserReply_User{
			Name:  svc.Name,
			Email: svc.Email,
		},
	}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.UserReply, error) {
	return &v1.UserReply{}, nil
}
