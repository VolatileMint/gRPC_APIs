/*
 * @Author: ShotaTakeda
 * @Date: 2023-04-10 15:52:14
 * @LastEditTime: 2023-04-24 10:54:02
 * @LastEditors: ShotaTakeda s-takeda@housei-inc.com
 * @Description:
 * @FilePath: \gRPC_APIs\endpoint\makeEndpoints.go
 */
 package endpoint

 import (
	 "gRPC_APIs/service"
 
	 "github.com/go-kit/log"
 
	 "github.com/go-kit/kit/endpoint"
 )
 
 type Endpoints struct {
	 AddUser       endpoint.Endpoint
	 ModUser       endpoint.Endpoint
	 DelUser       endpoint.Endpoint
	 DelUsersByOr  endpoint.Endpoint
	 ListUsers     endpoint.Endpoint
	 ListUsersByOr endpoint.Endpoint
 }
 
 func MakeEndpoints(svc service.Service, logger log.Logger) Endpoints {
	 var AddUser endpoint.Endpoint
	 {
		 AddUser = makeAddUserEndpoint(svc)
	 }
	 var ModUser endpoint.Endpoint
	 {
		 ModUser = makeModUserEndpoint(svc)
	 }
	 var DelUser endpoint.Endpoint
	 {
		 DelUser = makeDelUserEndpoint(svc)
	 }
	 var DelUsersByOr endpoint.Endpoint
	 {
		 DelUsersByOr = makeDelUsersByOrEndpoint(svc)
	 }
	 var ListUsers endpoint.Endpoint
	 {
		 ListUsers = makeListUsersEndpoint(svc)
	 }
	 var ListUsersByOr endpoint.Endpoint
	 {
		 ListUsersByOr = makeListUsersByOrEndpoint(svc)
	 }
	 return Endpoints{
		 AddUser,
		 ModUser,
		 DelUser,
		 DelUsersByOr,
		 ListUsers,
		 ListUsersByOr,
	 }
 }
 