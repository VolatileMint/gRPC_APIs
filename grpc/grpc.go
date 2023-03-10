package grpc

import (
	"gRPC_APIs/endpoint"
	testpb "gRPC_APIs/proto"

	"github.com/go-kit/log"

	"github.com/go-kit/kit/transport"
	gt "github.com/go-kit/kit/transport/grpc"
)

type gRPCServer struct {
	addUser   gt.Handler
	delUser   gt.Handler
	modUser   gt.Handler
	listUsers gt.Handler

	testpb.UnimplementedTestServiceServer
}

func NewGRPCServer(endpoints endpoint.Endpoints, logger log.Logger) testpb.TestServiceServer {
	options := []gt.ServerOption{
		gt.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
	}
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
		modUser: gt.NewServer(
			endpoints.ModUser,
			decodeModUserRequest,
			encodeModUserResponse,
		),
		listUsers: gt.NewServer(
			endpoints.ListUsers,
			decodeListUsersRequest,
			encodeListUsersResponse,
			append(options, gt.ServerBefore())...,
		),
	}
}

// grpcのチュートリアルより...
// func (s *gRPCServer) EchoName(ctx context.Context, req *testpb.Name) (*testpb.NameResp, error) {
// 	return &testpb.NameResp{
// 		Text: fmt.Sprintf("Hello, %s!", req.GetName()),
// 	}, nil
// }

// func (s *gRPCServer) ReturnFibo(ctx context.Context, req *testpb.ReturnFiboReq) (*testpb.ReturnFiboResp, error) {
// 	num := req.GetNum()
// 	// protoでint32なので、型をそれに合わせる
// 	fibo_array := []int32{1, 1}
// 	for i := int32(2); i < num; i++ {
// 		fibo_array = append(fibo_array, fibo_array[i-1]+fibo_array[i-2])
// 	}
// 	return &testpb.ReturnFiboResp{
// 		FiboArray: fibo_array,
// 	}, nil
// }
