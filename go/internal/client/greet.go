package client

import (
	"context"
	"time"

	"connectrpc.com/connect"
	pb "github.com/vinewz/clutchRPC/go/gen/clutch/v1"
	pbconnect "github.com/vinewz/clutchRPC/go/gen/clutch/v1/v1connect"
)

type GreetClient struct {
	Stub      pbconnect.GreetServiceClient
	TimeoutMS int
}

func (c *GreetClient) Greet(ctx context.Context, name string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(c.TimeoutMS)*time.Millisecond)
	defer cancel()

	req := connect.NewRequest(&pb.GreetRequest{Name: name})
	resp, err := c.Stub.Greet(ctx, req)
	if err != nil {
		return "", err
	}
	return resp.Msg.Greet, nil
}
