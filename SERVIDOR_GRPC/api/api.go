package api

import (
	"fmt"
	"net"

	proto "github.com/carlos/grpc/proto"
	"google.golang.org/grpc"
)

func Api_GRPC(address string) error {
	listen, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("error al escuchar: %v", err)
	}
	serve := grpc.NewServer()

	proto.RegisterJsonServer(serve, &Server{})

	fmt.Printf("Start HTTP/2 gRPC server on: %s \n", address)

	if err := serve.Serve(listen); err != nil {
		return fmt.Errorf("error al servir: %v", err)
	}

	return nil

}
