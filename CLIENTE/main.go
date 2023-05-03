package main

import (
	"context"
	"fmt"
	"log"

	proto "github.com/carlos/GRPCcliente/proto"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect :%s", err)
	}

	defer conn.Close()

	c := proto.NewGreeterClient(conn)

	response, err := c.SayHello(context.Background(), &proto.HelloRequest{Name: "Pruebas"})
	if err != nil {
		log.Fatalf("Error al llamar al método SayHello:%s", err)
	}
	fmt.Println(response.Message)
}
