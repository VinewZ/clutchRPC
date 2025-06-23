package server

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	pb "github.com/vinewz/clutchRPC/go/pb/clutch/v1"
)

type SayHiServiceServer struct {}

func (s *SayHiServiceServer) SayHi(ctx context.Context, req *connect.Request[pb.SayHiRequest]) (*connect.Response[pb.SayHiResponse], error) {
	fmt.Println(req.Msg.Greet)
	return &connect.Response[pb.SayHiResponse]{
		Msg: &pb.SayHiResponse{
			Greet: "Hello",
		},
	}, nil
}
