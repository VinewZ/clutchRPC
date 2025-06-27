package server

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	pb "github.com/vinewz/clutchRPC/go/gen/clutch/v1"
)

type GreetServiceServer struct{}

func (s *GreetServiceServer) Greet(ctx context.Context, req *connect.Request[pb.GreetRequest]) (*connect.Response[pb.GreetResponse], error) {
	return &connect.Response[pb.GreetResponse]{
		Msg: &pb.GreetResponse{
			Greet: fmt.Sprintf("Hello, %s!", req.Msg.Name),
		},
	}, nil
}
