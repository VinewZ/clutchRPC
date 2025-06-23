package server

import (
	"context"

	"connectrpc.com/connect"
	pb "github.com/vinewz/clutchRPC/go/pb/clutch/v1"
	"github.com/wailsapp/wails/v3/pkg/application"
)

type ToggleWindowServiceServer struct {
	App *application.App
}

func (s *ToggleWindowServiceServer) ToggleWindow(ctx context.Context, req *connect.Request[pb.ToggleWindowRequest]) (*connect.Response[pb.ToggleWindowResponse], error) {
	if req.Msg.IsVisible {
		s.App.Hide()
	} else {
		s.App.Show()
	}

	return &connect.Response[pb.ToggleWindowResponse]{
		Msg: &pb.ToggleWindowResponse{},
	}, nil
}
