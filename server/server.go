package server

import (
	"fmt"
	endpoints "gRPC_APIs/endpoint"
	grpctransport "gRPC_APIs/grpc"
	"gRPC_APIs/libs"
	testpb "gRPC_APIs/proto"
	"gRPC_APIs/service"
	deflog "log"
	"net"
	"os"
	"os/signal"

	"github.com/go-kit/log"
	"google.golang.org/grpc"
)

func Run() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}
	var (
		service    = service.NewSvc(logger)
		endpoints  = endpoints.MakeEndpoints(service, logger)
		grpcServer = grpctransport.NewGRPCServer(endpoints, logger)
	)
	// 8080番portのlistenerを作成
	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
	// grpcサーバーを作成
	s := grpc.NewServer()
	// サーバーにサービスを登録
	testpb.RegisterTestServiceServer(s, grpcServer)

	err = libs.Connect()
	if err != nil {
		panic(err)
	}
	// サーバーを稼働
	go func() {
		deflog.Printf("start gRPC server port: %v", port)
		s.Serve(listener)
	}()
	// Cntl+C が入力されたら止める
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	deflog.Println("stopping gRPC server...")
	s.GracefulStop()
}
