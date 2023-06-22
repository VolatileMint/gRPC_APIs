package grpc

import (
	"context"
	"fmt"
	"gRPC_APIs/model"
	pb "gRPC_APIs/proto"

	ts "google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

func decodeAddUserRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req, ok := grpcReq.(*pb.AddUserReq)
	if !ok {
		return nil, fmt.Errorf("grpc server decode request error")
	}
	request := &model.AddUserReq{
		User: model.TbTUser{
			Name:  req.User.Name,
			Email: req.User.Email,
		},
	}
	return request, nil
}
func encodeAddUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	res, ok := response.(*model.AddUserResp)
	if !ok {
		return nil, fmt.Errorf("grpc server encode response error")
	}
	created_at := ts.New(res.User.CreatedAt)
	updated_at := ts.New(res.User.UpdatedAt)
	resp := &pb.AddUserResp{
		User: &pb.User{
			Id:    int32(res.User.ID),
			Name:  res.User.Name,
			Email: res.User.Email,
			Timestamp: &pb.Timestamp{
				CreatedAt: created_at,
				UpdatedAt: updated_at,
			},
		},
	}
	return resp, nil
}
func (s *gRPCServer) AddUser(ctx context.Context, req *pb.AddUserReq) (*pb.AddUserResp, error) {
	_, resp, err := s.addUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.AddUserResp), nil
}

func decodeModUserRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req, ok := grpcReq.(*pb.ModUserReq)
	if !ok {
		return nil, fmt.Errorf("grpc server decode request error")
	}
	request := &model.ModUserReq{
		User: model.TbTUser{
			Model: gorm.Model{ID: uint(req.User.Id)},
			Name:  req.User.Name,
			Email: req.User.Email,
		},
	}
	return request, nil
}
func encodeModUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	res, ok := response.(*model.ModUserResp)
	if !ok {
		return nil, fmt.Errorf("grpc server encode response error")
	}
	created_at := ts.New(res.User.CreatedAt)
	updated_at := ts.New(res.User.UpdatedAt)
	resp := &pb.ModUserResp{
		User: &pb.User{
			Id:    int32(res.User.ID),
			Name:  res.User.Name,
			Email: res.User.Email,
			Timestamp: &pb.Timestamp{
				CreatedAt: created_at,
				UpdatedAt: updated_at,
			},
		},
	}
	return resp, nil
}
func (s *gRPCServer) ModUser(ctx context.Context, req *pb.ModUserReq) (*pb.ModUserResp, error) {
	_, resp, err := s.modUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ModUserResp), nil
}

func decodeDelUserRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req, ok := grpcReq.(*pb.DelUserReq)
	if !ok {
		return nil, fmt.Errorf("grpc server decode request error")
	}
	request := &model.DelUserReq{
		Ids: req.Ids,
	}
	return request, nil
}
func encodeDelUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	_, ok := response.(model.DelUserResp)
	if !ok {
		return nil, fmt.Errorf("grpc server encode response error")
	}
	return &pb.DelUserResp{}, nil
}
func (s *gRPCServer) DelUser(ctx context.Context, req *pb.DelUserReq) (*pb.DelUserResp, error) {
	_, resp, err := s.delUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.DelUserResp), nil
}

func decodeListUsersRequest(_ context.Context, request interface{}) (interface{}, error) {
	req, ok := request.(*pb.ListUsersReq)
	if !ok {
		return nil, fmt.Errorf("grpc server decode request error！")
	}
	var fields []*model.Field
	var key_value interface{}
	for _, v := range req.Fields {
		switch v.KeyValue.(type) {
		case *pb.Field_StrValue:
			key_value = v.GetStrValue()
		case *pb.Field_IntValue:
			key_value = v.GetIntValue()
		case *pb.Field_BoolValue:
			key_value = v.GetBoolValue()
		case *pb.Field_TimeValue:
			key_value = v.GetTimeValue()
		}
		fields = append(fields, &model.Field{
			Key:   v.Key,
			Value: key_value,
		})
	}
	return &model.ListUsersReq{
		Search: &model.Search{
			Fields: fields,
		},
	}, nil
}

func encodeListUsersResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp, ok := response.(*model.ListUsersResp)
	if !ok {
		return nil, fmt.Errorf("grpc server encode response error！")
	}
	Users := make([]*pb.User, len(resp.Users))
	for i, r := range resp.Users {
		created_at := ts.New(r.Model.CreatedAt)

		updated_at := ts.New(r.Model.UpdatedAt)
		Users[i] = &pb.User{
			Id:    int32(r.Model.ID),
			Email: r.Email,
			Name:  r.Name,
			Timestamp: &pb.Timestamp{
				CreatedAt: created_at,
				UpdatedAt: updated_at,
			},
		}

	}
	r := &pb.ListUsersResp{User: Users}
	return r, nil
}
func (s *gRPCServer) ListUsers(ctx context.Context, req *pb.ListUsersReq) (*pb.ListUsersResp, error) {
	_, resp, err := s.listUsers.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ListUsersResp), nil
}

func decodeListUsersByOrRequest(_ context.Context, request interface{}) (interface{}, error) {
	req, ok := request.(*pb.ListUsersByOrReq)
	if !ok {
		return nil, fmt.Errorf("grpc server decode request error！")
	}
	var fields []*model.Field
	var key_value interface{}
	for _, v := range req.Fields {
		switch v.KeyValue.(type) {
		case *pb.Field_StrValue:
			key_value = v.GetStrValue()
		case *pb.Field_IntValue:
			key_value = v.GetIntValue()
		case *pb.Field_BoolValue:
			key_value = v.GetBoolValue()
		case *pb.Field_TimeValue:
			key_value = v.GetTimeValue()
		}
		fields = append(fields, &model.Field{
			Key:   v.Key,
			Value: key_value,
		})
	}
	return &model.ListUsersByOrReq{
		Search: &model.Search{
			Fields:   fields,
			OrderBy:  req.OrderBy,
			Sort:     req.Sort,
			Page:     int(req.Page),
			PageSize: int(req.Pagesize),
		},
	}, nil
}

func encodeListUsersByOrResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp, ok := response.(*model.ListUsersByOrResp)
	if !ok {
		return nil, fmt.Errorf("grpc server encode response error！")
	}
	Users := make([]*pb.User, len(resp.Users))
	for i, r := range resp.Users {
		created_at := ts.New(r.Model.CreatedAt)

		updated_at := ts.New(r.Model.UpdatedAt)
		Users[i] = &pb.User{
			Id:    int32(r.Model.ID),
			Email: r.Email,
			Name:  r.Name,
			Timestamp: &pb.Timestamp{
				CreatedAt: created_at,
				UpdatedAt: updated_at,
			},
		}

	}
	r := &pb.ListUsersByOrResp{User: Users}
	return r, nil
}

func (s *gRPCServer) ListUsersByOr(ctx context.Context, req *pb.ListUsersByOrReq) (*pb.ListUsersByOrResp, error) {
	_, resp, err := s.listUsersByOr.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ListUsersByOrResp), nil
}

func decodeDelUsersByOrRequest(_ context.Context, request interface{}) (interface{}, error) {
	req, ok := request.(*pb.DelUsersByOrReq)
	if !ok {
		return nil, fmt.Errorf("grpc server decode request error！")
	}
	var fields []*model.Field
	var key_value interface{}
	for _, v := range req.Fields {
		switch v.KeyValue.(type) {
		case *pb.Field_StrValue:
			key_value = v.GetStrValue()
		case *pb.Field_IntValue:
			key_value = v.GetIntValue()
		case *pb.Field_BoolValue:
			key_value = v.GetBoolValue()
		case *pb.Field_TimeValue:
			key_value = v.GetTimeValue()
		}
		fields = append(fields, &model.Field{
			Key:   v.Key,
			Value: key_value,
		})
	}
	return &model.ListUsersByOrReq{
		Search: &model.Search{
			Fields: fields,
		},
	}, nil
}

func encodeDelUsersByOrResponse(_ context.Context, response interface{}) (interface{}, error) {
	_, ok := response.(*model.DelUsersByOrResp)
	if !ok {
		return nil, fmt.Errorf("grpc server encode response error")
	}
	return &pb.DelUserResp{}, nil
}

func (s *gRPCServer) DelUsersByOr(ctx context.Context, req *pb.DelUsersByOrReq) (*pb.DelUsersByOrResp, error) {
	_, resp, err := s.delUsersByOr.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.DelUsersByOrResp), nil
}
