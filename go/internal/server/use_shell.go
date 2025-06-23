package server

import (
	"context"
	"fmt"
	"os/exec"
	"sync"
	"time"

	"connectrpc.com/connect"
	v1 "github.com/vinewz/clutchRPC/go/pb/clutch/v1"
	"github.com/wailsapp/wails/v3/pkg/application"
)

type UseShellServiceServer struct {
	App *application.App

	mu        sync.Mutex
	confirmCh chan bool
}

func (s *UseShellServiceServer) UseShell(ctx context.Context, req *connect.Request[v1.UseShellRequest]) (*connect.Response[v1.UseShellResponse], error) {
	s.mu.Lock()
	if s.confirmCh != nil {
		s.mu.Unlock()
		return nil, fmt.Errorf("another confirmation is pending")
	}
	ch := make(chan bool, 1)
	s.confirmCh = ch
	cmd := req.Msg.Command
	s.mu.Unlock()

	s.App.EmitEvent("clutch:require-confirmation", map[string]string{
		"appName": req.Msg.AppName,
		"command": cmd,
	})

	fmt.Printf("Waiting for confirmation to run command: %q\n", cmd)

	select {
	case confirmed := <-ch:
		if !confirmed {
			return nil, fmt.Errorf("user declined %q", cmd)
		}
	case <-time.After(30 * time.Second):
		s.mu.Lock()
		s.confirmCh = nil
		s.mu.Unlock()
		return nil, fmt.Errorf("confirmation timed out")
	}

	s.mu.Lock()
	s.confirmCh = nil
	s.mu.Unlock()

	out, err := exec.Command("sh", "-c", cmd).CombinedOutput()
	return &connect.Response[v1.UseShellResponse]{
		Msg: &v1.UseShellResponse{
			Output: string(out),
		},
	}, err
}
