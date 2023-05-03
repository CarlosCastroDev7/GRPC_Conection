package api

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	proto "github.com/Carlos/grpc-rest/proto"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func Api_GRPC(address string) error {
	listen, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("error al escuchar: %v", err)
	}
	serve := grpc.NewServer()

	proto.RegisterPruebasgrpcServer(serve, &Server{})

	fmt.Printf("Start HTTP/2 gRPC server on: %s \n", address)

	if err := serve.Serve(listen); err != nil {
		return fmt.Errorf("error al servir: %v", err)
	}

	return nil
}

func Api_REST(address string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := proto.RegisterPruebasgrpcHandlerFromEndpoint(ctx, mux, "127.0.0.1:50001", opts)
	if err != nil {
		fmt.Printf("could not register service Ping: %s", err)
	}

	log.Printf("starting HTTP/1.1 with TLS REST server on %s", address)
	http.ListenAndServe(address, mux)
	return nil
}
