package grpc

import (
	"gRPC_APIs/endpoint"
	testpb "gRPC_APIs/proto"

	"github.com/go-kit/log"

	gt "github.com/go-kit/kit/transport/grpc"
)

type gRPCServer struct {
	addUser       gt.Handler
	delUser       gt.Handler
	delUsersByOr  gt.Handler
	modUser       gt.Handler
	listUsers     gt.Handler
	listUsersByOr gt.Handler

	testpb.UnimplementedTestServiceServer
}

func NewGRPCServer(endpoints endpoint.Endpoints, logger log.Logger) testpb.TestServiceServer {
	// 使わなくても動く、要検証
	// options := []gt.ServerOption{
	// 	gt.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
	// }
	return &gRPCServer{
		addUser: gt.NewServer(
			endpoints.AddUser,
			decodeAddUserRequest,
			encodeAddUserResponse,
		),
		delUser: gt.NewServer(
			endpoints.DelUser,
			decodeDelUserRequest,
			encodeDelUserResponse,
		),
		delUsersByOr: gt.NewServer(
			endpoints.DelUsersByOr,
			decodeDelUsersByOrRequest,
			encodeDelUsersByOrResponse,
		),
		modUser: gt.NewServer(
			endpoints.ModUser,
			decodeModUserRequest,
			encodeModUserResponse,
		),
		listUsers: gt.NewServer(
			endpoints.ListUsers,
			decodeListUsersRequest,
			encodeListUsersResponse,
			// append(options, gt.ServerBefore())...,
		),
		listUsersByOr: gt.NewServer(
			endpoints.ListUsersByOr,
			decodeListUsersByOrRequest,
			encodeListUsersByOrResponse,
		),
	}
}
