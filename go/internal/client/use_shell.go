package client

import (
	"context"
	"time"

	"connectrpc.com/connect"
	pb "github.com/vinewz/clutchRPC/go/pb/clutch/v1"
	pbconnect "github.com/vinewz/clutchRPC/go/pb/clutch/v1/v1connect"
)

type UseShellClient struct {
	Stub      pbconnect.UseShellServiceClient
	TimeoutMS int
}

func (c *UseShellClient) UseShell(ctx context.Context, cmd string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(c.TimeoutMS)*time.Millisecond)
	defer cancel()

	req := connect.NewRequest(&pb.UseShellRequest{Command: cmd})
	resp, err := c.Stub.UseShell(ctx, req)
	if err != nil {
		return "", err
	}
	return resp.Msg.GetOutput(), nil
}
