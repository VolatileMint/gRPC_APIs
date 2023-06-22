package model

import (
	"fmt"
	"gRPC_APIs/libs"
	"reflect"

	"gorm.io/gorm"
)

func GetAll(model interface{}, s *Search) *gorm.DB {
	db := libs.DB.Model(model)
	sort := "desc"
	orderBy := "created_at"
	if len(s.Sort) > 0 {
		sort = s.Sort
	}
	if len(s.OrderBy) > 0 {
		orderBy = s.OrderBy
	}
	db = db.Order(fmt.Sprintf("%s %s", orderBy, sort))
	db.Scopes(FoundByOr(s.Fields))
	return db
}

func Pagenate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize < 0:
			pageSize = -1
		case pageSize == 0:
			pageSize = 0
		}
		offset := (page - 1) * pageSize
		if page < 0 {
			offset = -1
		}
		return db.Offset(offset).Limit(pageSize)
	}
}

func GetAllByOr(model interface{}, s *Search) *gorm.DB {
	db := libs.DB.Model(model)
	sort := "desc"
	orderBy := "created_at"
	if len(s.Sort) > 0 {
		sort = s.Sort
	}
	if len(s.OrderBy) > 0 {
		orderBy = s.OrderBy
	}

	db = db.Order(fmt.Sprintf("%s %s", orderBy, sort))

	db.Scopes(FoundByOr(s.Fields))

	return db
}

func FoundByOr(fields []*Field) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(fields) > 0 {
			for _, field := range fields {
				if field != nil {
					if field.Condition == "" {
						field.Condition = "="
					}
					if value, ok := field.Value.(int); ok {
						if value > 0 {
							db = db.Or(fmt.Sprintf("%s %s ?", field.Key, field.Condition), value)
						}
					} else if value, ok := field.Value.(uint); ok {
						if value > 0 {
							db = db.Or(fmt.Sprintf("%s %s ?", field.Key, field.Condition), value)
						}
					} else if value, ok := field.Value.(int32); ok {
						if value > 0 {
							db = db.Or(fmt.Sprintf("%s %s ?", field.Key, field.Condition), value)
						}
					} else if value, ok := field.Value.(string); ok {
						if len(value) > 0 {
							db = db.Or(fmt.Sprintf("%s %s ?", field.Key, field.Condition), value)
						}
					} else if value, ok := field.Value.([]int); ok {
						if len(value) > 0 {
							db = db.Or(fmt.Sprintf("%s %s ?", field.Key, field.Condition), value)
						}
					} else if value, ok := field.Value.([]string); ok {
						if len(value) > 0 {
							db = db.Or(fmt.Sprintf("%s %s ?", field.Key, field.Condition), value)
						}
					} else {
						fmt.Printf("未知数据类型：%+v", field.Value)
					}
				}
			}
		}
		return db
	}
}

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize < 0:
			pageSize = -1
		case pageSize == 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		if page < 0 {
			offset = -1
		}
		return db.Offset(offset).Limit(pageSize)
	}
}

// 削除したいテーブルの構造体と削除条件をもらい削除する
func DelByOr(value interface{}, s *Search) (*gorm.DB, int64, error) {
	newStruct := reflect.New(reflect.TypeOf(value)).Elem()

	var instance interface{}
	name := reflect.TypeOf(value)
	instance_type := name.Name()
	fmt.Print("this is ", instance_type)

	switch instance_type {
	case "User":
		instance = newStruct.Interface().(User)
	}
	// get all the data with the struct and search
	all := GetAllByOr(instance, s)
	var count int64

	if err := all.Debug().Count(&count).Error; err != nil {
		return nil, 0, err
	}

	all = all.Scopes(Paginate(s.Page, s.PageSize))

	// delete the row/rows of data
	resp := all.Debug().Delete(&instance)

	return resp, count, nil
}
