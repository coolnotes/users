package users

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/coolnotes/users/src/rpc"
	"github.com/coolnotes/users/src/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type App struct {
}

func (a *App) Start() {
	fmt.Println("Running!")
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	// ...
	grpcServer := grpc.NewServer(opts...)
	rpc.RegisterUserServer(grpcServer, server.NewUserServer())
	reflection.Register(grpcServer)
	// pb.RegisterRouteGuideServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
