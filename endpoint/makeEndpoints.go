package endpoint

import (
	"gRPC_APIs/service"

	"github.com/go-kit/log"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	AddUser   endpoint.Endpoint
	ModUser   endpoint.Endpoint
	DelUser   endpoint.Endpoint
	ListUsers endpoint.Endpoint
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
	var ListUsers endpoint.Endpoint
	{
		ListUsers = makeListUsersEndpoint(svc)
	}
	return Endpoints{
		AddUser,
		ModUser,
		DelUser,
		ListUsers,
	}
}
