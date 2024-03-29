package endpoint

import (
	"context"
	"fmt"
	"gRPC_APIs/model"
	service "gRPC_APIs/service"

	"github.com/go-kit/kit/endpoint"
)

// API step add endpoint
func makeAddUserEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*model.AddUserReq)
		if !ok {
			return nil, fmt.Errorf("err:AddUserEndpoint")
		}
		resp, err := s.AddUser(ctx, req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}
func makeModUserEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*model.ModUserReq)
		if !ok {
			return nil, fmt.Errorf("err:ModUserEndpoint")
		}
		resp, err := s.ModUser(ctx, req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}
func makeDelUserEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*model.DelUserReq)
		if !ok {
			return nil, fmt.Errorf("err:DelUserEndpoint")
		}
		_, err = s.DelUser(ctx, req)
		if err != nil {
			return nil, err
		}
		return model.DelUserResp{}, nil
	}
}
func makeDelUsersByOrEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*model.DelUsersByOrReq)
		if !ok {
			return nil, fmt.Errorf("err:DelUserEndpoint")
		}
		_, err = s.DelUsersByOr(ctx, req)
		if err != nil {
			return nil, err
		}
		return model.DelUsersByOrResp{}, nil
	}
}
func makeListUsersEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*model.ListUsersReq)
		if !ok {
			return nil, fmt.Errorf("err:ListUsersEndpoint")
		}
		resp, err := s.ListUsers(ctx, req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

func makeListUsersByOrEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*model.ListUsersByOrReq)
		if !ok {
			return nil, fmt.Errorf("err:ListUsersByOrEndpoint")
		}
		resp, err := s.ListUsersByOr(ctx, req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}
