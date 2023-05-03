package api

import (
	"context"
	"fmt"

	proto "github.com/Carlos/grpc-rest/proto"
)

type Server struct {
	proto.UnimplementedPruebasgrpcServer
}

func (*Server) Saludar(ctx context.Context, r *proto.Request) (*proto.Response, error) {
	return &proto.Response{Message: fmt.Sprintf("El valor tiene parametros %v", r.Message)}, nil
}
