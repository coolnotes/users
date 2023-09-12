package server

import (
	"context"

	"github.com/coolnotes/users/src/rpc"
)

type Server struct {
	rpc.UnimplementedUserServer
}

func NewUserServer() rpc.UserServer {
	return &Server{}
}

func (Server) Create(context.Context, *rpc.UserCreationRequest) (*rpc.UserResponse, error) {
	return &rpc.UserResponse{
		Id:    1,
		Email: "voff.web@gmail.com",
	}, nil
}
