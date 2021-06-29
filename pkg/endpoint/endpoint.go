package endpoint

import (
	"context"

	"github.com/Zubayear/bruce-almighty/pkg/service"
	"github.com/go-kit/kit/endpoint"
)

type CreateUserRequest struct {
	Name string `json:"artist_name"`
}

type GetUserRequest struct {
	Id string `json:"id"`
}

type CreateUserResponse struct {
	Name string `json:"artist_name"`
	Err  string `json:"err"`
}
type GetUserResponse struct {
	Name string `json:"artist_name"`
	Err  string `json:"err"`
}

func MakeCreateUserEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserRequest)
		s, err := svc.CreateUser(ctx, req.Name)
		if err != nil {
			return CreateUserResponse{Name: s, Err: err.Error()}, nil
		}
		return CreateUserResponse{Name: s, Err: ""}, nil
	}
}

func MakeGetUserEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserRequest)
		s, err := svc.GetUser(ctx, req.Id)
		if err != nil {
			return GetUserResponse{Name: s, Err: err.Error()}, nil
		}
		return GetUserResponse{Name: s, Err: ""}, nil
	}
}

type Endpoints struct {
	CreateUser endpoint.Endpoint
	GetUser    endpoint.Endpoint
}

func MakeEndpoints(s service.Service) Endpoints {
	return Endpoints{
		CreateUser: MakeCreateUserEndpoint(s),
		GetUser:    MakeGetUserEndpoint(s),
	}
}
