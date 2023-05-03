package api

import (
	"context"
	"fmt"

	proto "github.com/carlos/grpc/proto"
)

type Server struct {
	proto.UnimplementedJsonServer
}

func (*Server) SayHello(ctx context.Context, r *proto.JsonRequest) (*proto.JsonResponse, error) {
	return &proto.JsonResponse{Json: fmt.Sprintf("El valor tiene parametros %v", r.Params)}, nil
}
