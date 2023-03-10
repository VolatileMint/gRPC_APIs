package service

import (
	"context"
	"fmt"
	"gRPC_APIs/libs"
	"gRPC_APIs/model"

	"gorm.io/gorm"
)

func (s UserService) AddUser(ctx context.Context, req *model.AddUserReq) (*model.AddUserResp, error) {
	newuser := model.TbTUser{
		Name:  req.User.Name,
		Email: req.User.Email,
	}
	err := libs.DB.Create(&newuser).Debug().Error
	if err != nil {
		return nil, err
	}
	response := &model.AddUserResp{
		User: model.TbTUser{
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
	moduser := model.TbTUser{
		Model: gorm.Model{ID: req.User.Model.ID},
		Name:  req.User.Name,
		Email: req.User.Email,
	}
	err := libs.DB.Updates(&moduser).Debug().Error
	if err != nil {
		return nil, err
	}

	response := &model.ModUserResp{
		User: model.TbTUser{
			Name:  moduser.Name,
			Email: moduser.Email,
			Model: gorm.Model{
				ID: moduser.ID,
				// DBに接続できていないので仮
				CreatedAt: moduser.CreatedAt,
				DeletedAt: moduser.DeletedAt,
			},
		},
	}
	return response, nil
}
func (s UserService) DelUser(ctx context.Context, req *model.DelUserReq) error {
	for _, rid := range req.Ids {
		user := model.TbTUser{Model: gorm.Model{ID: uint(rid)}}
		// 削除したいユーザーが存在するかの確認
		err := libs.DB.First(&user)
		if err != nil {
			return fmt.Errorf("ユーザーが見つかりませんでした%v", err)
		}

		resperr := libs.DB.Delete(&user)
		if resperr != nil {
			return fmt.Errorf("削除に失敗しました%v", err)
		}
	}
	return nil
}
func (s UserService) ListUsers(ctx context.Context, req *model.ListUsersReq) (*model.ListUsersResp, error) {
	resp, _, err := model.GetAllUser(req.Search)
	if err != nil {
		return nil, err
	}
	response := &model.ListUsersResp{Users: resp}
	return response, nil
}