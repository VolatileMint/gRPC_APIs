package service

import (
	"context"
	"fmt"
	"gRPC_APIs/libs"
	"gRPC_APIs/model"

	"gorm.io/gorm"
)

func (s UserService) AddUser(ctx context.Context, req *model.AddUserReq) (*model.AddUserResp, error) {
	newuser := model.User{
		Name:  req.UserReq.Name,
		Email: req.UserReq.Email,
	}
	err := libs.DB.Create(&newuser).Debug().Error
	if err != nil {
		return nil, err
	}
	response := &model.AddUserResp{
		UserReq: model.User{
			Name:  newuser.Name,
			Email: newuser.Email,
			Model: gorm.Model{
				ID:        newuser.ID,
				CreatedAt: newuser.CreatedAt,
				DeletedAt: newuser.DeletedAt,
			},
		},
	}
	return response, nil
}
func (s UserService) ModUser(ctx context.Context, req *model.ModUserReq) (*model.ModUserResp, error) {
	moduser := model.User{
		Model: gorm.Model{ID: req.UserReq.Model.ID},
		Name:  req.UserReq.Name,
		Email: req.UserReq.Email,
	}
	err := libs.DB.Updates(&moduser).Debug().Error
	if err != nil {
		return nil, err
	}

	response := &model.ModUserResp{
		UserReq: model.User{
			Name:  moduser.Name,
			Email: moduser.Email,
			Model: gorm.Model{
				ID:        moduser.ID,
				CreatedAt: moduser.CreatedAt,
				DeletedAt: moduser.DeletedAt,
			},
		},
	}
	return response, nil
}

func (s UserService) DelUser(ctx context.Context, req *model.DelUserReq) (*model.DelUserResp, error) {
	for _, rid := range req.Ids {
		user := model.User{Model: gorm.Model{ID: uint(rid)}}
		// 削除したいユーザーが存在するかの確認
		err := libs.DB.First(&user).Debug().Error
		if err != nil {
			return nil, fmt.Errorf("ユーザーが見つかりませんでした%v", err)
		}

		resperr := libs.DB.Delete(&user)
		if resperr != nil {
			return nil, fmt.Errorf("削除に失敗しました%v", resperr)
		}
	}
	return &model.DelUserResp{}, nil
}

func (s UserService) ListUsers(ctx context.Context, req *model.ListUsersReq) (*model.ListUsersResp, error) {
	resp, _, err := model.GetAllUser(req.Search)
	if err != nil {
		return nil, err
	}
	response := &model.ListUsersResp{Users: resp}
	return response, nil
}

func (s UserService) ListUsersByOr(ctx context.Context, req *model.ListUsersByOrReq) (*model.ListUsersByOrResp, error) {
	resp, _, err := model.GetAllUserByOr(req.Search)
	if err != nil {
		return nil, err
	}
	response := &model.ListUsersByOrResp{Users: resp}
	return response, nil
}

func (s UserService) DelUsersByOr(ctx context.Context, req *model.DelUsersByOrReq) (*model.DelUsersByOrResp, error) {
	tb_t_users_list := model.User{}
	resp, _, err := model.DelByOr(tb_t_users_list, req.Search)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, resp.Error
	}
	return &model.DelUsersByOrResp{}, nil
}
