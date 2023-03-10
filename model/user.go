package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type AddUserReq struct {
	User TbTUser
}
type AddUserResp struct {
	User TbTUser
}
type ModUserReq struct {
	User TbTUser
}
type ModUserResp struct {
	User TbTUser
}
type DelUserReq struct {
	Ids []int32
}

type Empty struct {
}

type ListUsersReq struct {
	Search *Search
}
type ListUsersResp struct {
	Users []*TbTUser
}
type Search struct {
	Fields  []*Field `json:"fields"`
	OrderBy string   `json:"order_by"`
	Sort    string   `json:"sort"`
	Limit   int      `json:"limit"`
	Offset  int      `json:"offset"`
}

type Field struct {
	Condition string      `json:"condition"`
	Key       string      `json:"key"`
	Value     interface{} `json:"value"`
}

type TbTUser struct {
	gorm.Model
	Email string
	Name  string
}

func GetAllUser(s *Search) ([]*TbTUser, int64, error) {
	// 元が理解できなかった
	var Users []*TbTUser
	var count int64

	for i := 0; i < 3; i++ {
		time := time.Now()
		new := &TbTUser{
			Model: gorm.Model{
				ID:        uint(i),
				CreatedAt: time,
				UpdatedAt: time,
			},
			Name:  fmt.Sprintf("test%v", i),
			Email: fmt.Sprintf("test%v@test.com", i),
		}
		Users = append(Users, new)
		count++
	}
	return Users, count, nil
}
