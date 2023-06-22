/*
 * @Author: ShotaTakeda
 * @Date: 2023-04-10 15:52:14
 * @LastEditTime: 2023-04-24 10:59:31
 * @LastEditors: ShotaTakeda s-takeda@housei-inc.com
 * @Description:
 * @FilePath: \gRPC_APIs\service\newUser.go
 */
 package service

 import (
	 "context"
	 "gRPC_APIs/model"
 
	 "github.com/go-kit/log"
 )
 
 type Service interface {
	 AddUser(ctx context.Context, u *model.AddUserReq) (*model.AddUserResp, error)
	 ModUser(ctx context.Context, u *model.ModUserReq) (*model.ModUserResp, error)
	 DelUser(ctx context.Context, u *model.DelUserReq) (*model.DelUserResp, error)
	 DelUsersByOr(ctx context.Context, u *model.DelUsersByOrReq) (*model.DelUsersByOrResp, error)
	 ListUsers(ctx context.Context, req *model.ListUsersReq) (*model.ListUsersResp, error)
	 ListUsersByOr(ctx context.Context, req *model.ListUsersByOrReq) (*model.ListUsersByOrResp, error)
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
 