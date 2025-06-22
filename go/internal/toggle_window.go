package internal

import (
	"context"

	"connectrpc.com/connect"
	pb "github.com/vinewz/clutchRPC/go/pb/clutch/v1"
	"github.com/wailsapp/wails/v3/pkg/application"
)

type ToggleWindowServiceServer struct {
	App       *application.App
	IsVisible bool
}

func (s *ToggleWindowServiceServer) ToggleWindow(ctx context.Context, req *connect.Request[pb.ToggleWindowRequest]) (*connect.Response[pb.ToggleWindowResponse], error) {
	if s.IsVisible {
		s.App.Hide()
		s.IsVisible = false
	} else {
		s.App.Show()
		s.IsVisible = true
	}
	return &connect.Response[pb.ToggleWindowResponse]{
		Msg: &pb.ToggleWindowResponse{},
	}, nil
}
