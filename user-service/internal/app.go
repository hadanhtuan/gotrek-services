package internal

import (
	"fmt"
	"log"
	"net"
	api "user-service/api"
	userProto "user-service/proto/user"

	pkg "github.com/hadanhtuan/go-sdk"
	"google.golang.org/grpc"
)

func InitGRPCServer(app *pkg.App) error {
	userServiceHost := fmt.Sprintf(
		"%s:%s",
		app.Config.GRPC.UserServiceHost,
		app.Config.GRPC.UserServicePort,
	)
	lis, err := net.Listen("tcp", userServiceHost)
	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}

	s := grpc.NewServer()
	userProto.RegisterUserServiceServer(s, &api.UserController{})

	log.Printf("gRPC Server started on %s", userServiceHost)

	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}

	fmt.Println("Server down")

	return nil
}