package service

import (
	"context"
	"gRPC_APIs/model"

	"github.com/go-kit/log"
)

type Service interface {
	// tb_t_notice table's API
	AddUser(ctx context.Context, u *model.AddUserReq) (*model.AddUserResp, error)
	ModUser(ctx context.Context, u *model.ModUserReq) (*model.ModUserResp, error)
	DelUser(ctx context.Context, u *model.DelUserReq) error
	ListUsers(ctx context.Context, req *model.ListUsersReq) (*model.ListUsersResp, error)
}

func NewSvc(logger log.Logger) Service {
	var svc Service
	{
		svc = NewUserService()
	}
	return svc
}

func NewUserService() Service {
	return UserService{}
}

type UserService struct{}
