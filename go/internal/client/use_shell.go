package client

import (
	"context"
	"time"

	"connectrpc.com/connect"
	pb "github.com/vinewz/clutchRPC/go/gen/clutch/v1"
	pbconnect "github.com/vinewz/clutchRPC/go/gen/clutch/v1/v1connect"
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
	return resp.Msg.Output, nil
}

func (c *UseShellClient) ConfirmShell(ctx context.Context, allow bool) error {
  // same timeout pattern
  ctx, cancel := context.WithTimeout(ctx, time.Duration(c.TimeoutMS)*time.Millisecond)
  defer cancel()

  req := connect.NewRequest(&pb.ConfirmShellRequest{
    Allow: allow,
  })
  // we don’t care about the empty response message
  if _, err := c.Stub.ConfirmShell(ctx, req); err != nil {
    return err
  }
  return nil
}
