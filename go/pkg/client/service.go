package clutchrpc

import (
	"context"
	"fmt"
	"net/http"

	"github.com/vinewz/clutchRPC/go/internal/client"
	pbconnect "github.com/vinewz/clutchRPC/go/pb/clutch/v1/v1connect"
)

type Client struct {
	*client.SayHiClient
	*client.UseShellClient
	*client.ToggleWindowClient
}

func New(ctx context.Context, port int, timeoutMS int) (*Client, error) {
	httpClient := &http.Client{}
	baseURL := fmt.Sprintf("http://localhost:%d", port)

	sayHiStub := pbconnect.NewSayHiServiceClient(httpClient, baseURL)
	shellStub := pbconnect.NewUseShellServiceClient(httpClient, baseURL)
	toggleStub := pbconnect.NewToggleWindowServiceClient(httpClient, baseURL)

	return &Client{
		SayHiClient:        &client.SayHiClient{Stub: sayHiStub, TimeoutMS: timeoutMS},
		UseShellClient:     &client.UseShellClient{Stub: shellStub, TimeoutMS: timeoutMS},
		ToggleWindowClient: &client.ToggleWindowClient{Stub: toggleStub, TimeoutMS: timeoutMS},
	}, nil
}
