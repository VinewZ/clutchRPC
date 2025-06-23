package client

import (
	"context"
	"time"

	"connectrpc.com/connect"
	pb "github.com/vinewz/clutchRPC/go/pb/clutch/v1"
	pbconnect "github.com/vinewz/clutchRPC/go/pb/clutch/v1/v1connect"
)

type SayHiClient struct {
	Stub       pbconnect.SayHiServiceClient
	TimeoutMS int
}

func (c *SayHiClient) SayHi(ctx context.Context, greet string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(c.TimeoutMS)*time.Millisecond)
	defer cancel()

	req := connect.NewRequest(&pb.SayHiRequest{Greet: greet})
	resp, err := c.Stub.SayHi(ctx, req)
	if err != nil {
		return "", err
	}
	return resp.Msg.Greet, nil
}
