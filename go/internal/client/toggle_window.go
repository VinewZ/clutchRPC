package client

import (
	"context"
	"time"

	"connectrpc.com/connect"
	pb "github.com/vinewz/clutchRPC/go/pb/clutch/v1"
	pbconnect "github.com/vinewz/clutchRPC/go/pb/clutch/v1/v1connect"
)

type ToggleWindowClient struct {
	Stub      pbconnect.ToggleWindowServiceClient
	TimeoutMS int
}

func (c *ToggleWindowClient) ToggleWindow(ctx context.Context, isVisible bool) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(c.TimeoutMS)*time.Millisecond)
	defer cancel()

	req := connect.NewRequest(&pb.ToggleWindowRequest{
		IsVisible: isVisible,
	})
	_, err := c.Stub.ToggleWindow(ctx, req)
	return err
}
