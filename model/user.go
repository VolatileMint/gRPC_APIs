/*
 * @Author: ShotaTakeda
 * @Date: 2023-04-10 15:52:14
 * @LastEditTime: 2023-04-24 11:01:14
 * @LastEditors: ShotaTakeda s-takeda@housei-inc.com
 * @Description:
 * @FilePath: \gRPC_APIs\model\user.go
 */
package model

import (
	"gorm.io/gorm"
)

type AddUserReq struct {
	UserReq User
}
type AddUserResp struct {
	UserReq User
}

type ModUserReq struct {
	UserReq User
}
type ModUserResp struct {
	UserReq User
}

type DelUserReq struct {
	Ids []int32
}
type DelUserResp struct {
}

type ListUsersReq struct {
	Search *Search
}
type ListUsersResp struct {
	Users []*User
}

type ListUsersByOrReq struct {
	Search *Search
}
type ListUsersByOrResp struct {
	Users []*User
}

type DelUsersByOrReq struct {
	Search *Search
}
type DelUsersByOrResp struct {
}

type Search struct {
	Fields   []*Field `json:"fields"`
	OrderBy  string   `json:"order_by"`
	Sort     string   `json:"sort"`
	Page     int      `json:"page"`
	PageSize int      `json:"page_size"`
}

type Field struct {
	Condition string      `json:"condition"`
	Key       string      `json:"key"`
	Value     interface{} `json:"value"`
}

type User struct {
	gorm.Model
	Email string
	Name  string
}

func GetAllUser(s *Search) ([]*User, int64, error) {
	var Users []*User
	var count int64

	for _, f := range s.Fields {
		f.Key = "tb_t_users." + f.Key
	}
	all := GetAll(&User{}, s)
	if err := all.Debug().Count(&count).Error; err != nil {
		return nil, 0, err
	}

	all = all.Scopes(Paginate(s.Page, s.PageSize))
	if err := all.Debug().Find(&Users).Error; err != nil {
		return nil, 0, err
	}
	return Users, count, nil
}
func GetAllUserByOr(s *Search) ([]*User, int64, error) {
	var Users []*User
	var count int64

	for _, f := range s.Fields {
		f.Key = "tb_t_users." + f.Key
	}
	all := GetAllByOr(&User{}, s)
	if err := all.Debug().Count(&count).Error; err != nil {
		return nil, 0, err
	}

	all = all.Scopes(Paginate(s.Page, s.PageSize))

	if err := all.Debug().Find(&Users).Error; err != nil {
		return nil, 0, err
	}
	return Users, count, nil
}
