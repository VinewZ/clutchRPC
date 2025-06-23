// go/client/client.go
package client

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"time"

	"connectrpc.com/connect"
	"golang.org/x/net/http2"

	pb "github.com/vinewz/clutchRPC/go/pb/clutch/v1"
	pbconnect "github.com/vinewz/clutchRPC/go/pb/clutch/v1/v1connect"
)

// Client holds one field per RPC service…
type Client struct {
	SayHi    pbconnect.SayHiServiceClient
	UseShell pbconnect.UseShellServiceClient
	Toggle   pbconnect.ToggleWindowServiceClient
}

// New creates a Client talking to your server at baseURL (e.g. "http://localhost:9023").
// It uses HTTP/2 clear-text (h2c) under the hood but falls back to HTTP/1.1 if needed.
func New(baseURL string, timeout time.Duration) *Client {
	// 1. Build an http.Client that supports h2c:
	h2cTr := &http2.Transport{
		AllowHTTP: true,
		// DialTLS is only used for HTTP/2 over clear-text:
		DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
			return net.Dial(network, addr)
		},
	}
	httpClient := &http.Client{
		Transport: h2cTr,
		Timeout:   timeout,
	}

	// 2. Now create each generated Connect client:
	return &Client{
		SayHi:    pbconnect.NewSayHiServiceClient(httpClient, baseURL),
		UseShell: pbconnect.NewUseShellServiceClient(httpClient, baseURL),
		Toggle:   pbconnect.NewToggleWindowServiceClient(httpClient, baseURL),
	}
}

// Convenience wrappers:

func (c *Client) SayHiCtx(ctx context.Context, name string) (*pb.SayHiResponse, error) {
	req := connect.NewRequest(&pb.SayHiRequest{ /* check your .proto for the exact field name: maybe "name" or "who" */ })
	resp, err := c.SayHi.SayHi(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.Msg, nil
}

func (c *Client) UseShellCtx(ctx context.Context, app, cmd string) (*pb.UseShellResponse, error) {
	req := connect.NewRequest(&pb.UseShellRequest{AppName: app, Command: cmd})
	resp, err := c.UseShell.UseShell(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.Msg, nil
}

func (c *Client) ToggleWindowCtx(ctx context.Context) (*pb.ToggleWindowResponse, error) {
	req := connect.NewRequest(&pb.ToggleWindowRequest{})
	resp, err := c.Toggle.ToggleWindow(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.Msg, nil
}
