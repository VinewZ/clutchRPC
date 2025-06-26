package server

import (
	"context"
	"fmt"
	"os/exec"
	"sync"
	"time"

	"connectrpc.com/connect"
	v1 "github.com/vinewz/clutchRPC/go/gen/clutch/v1"
	"github.com/wailsapp/wails/v3/pkg/application"
)

type UseShellServiceServer struct {
	App *application.App

	mu        sync.Mutex
	confirmCh chan bool
}

// UseShell sends a confirmation event to the frontend and waits (with timeout) for the user's answer.
// It then runs `sh -c <command>` and returns its combined output or an error if the command failed/cancelled.
func (s *UseShellServiceServer) UseShell(
	ctx context.Context,
	req *connect.Request[v1.UseShellRequest],
) (*connect.Response[v1.UseShellResponse], error) {
	s.mu.Lock()
	if s.confirmCh != nil {
		s.mu.Unlock()
		return nil, fmt.Errorf("Another confirmation is pending")
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
		s.mu.Lock()
		s.confirmCh = nil
		s.mu.Unlock()

		if !confirmed {
			return nil, fmt.Errorf("Command %q cancelled by user", cmd)
		}

	case <-time.After(30 * time.Second):
		s.mu.Lock()
		s.confirmCh = nil
		s.mu.Unlock()
		return nil, fmt.Errorf("Confirmation for %q timed out", cmd)
	}

	outBytes, err := exec.Command("sh", "-c", cmd).CombinedOutput()
	output := string(outBytes)
	if err != nil {
		return nil, fmt.Errorf("Command %q failed: %w\noutput:\n%s", cmd, err, output)
	}

	return connect.NewResponse(&v1.UseShellResponse{
		Output: output,
	}), nil
}

// ConfirmShell is called by the frontend to send back the user's decision.
// It unblocks the pending UseShell call by sending `allow` into the channel.
func (s *UseShellServiceServer) ConfirmShell(
	ctx context.Context,
	req *connect.Request[v1.ConfirmShellRequest],
) (*connect.Response[v1.ConfirmShellResponse], error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.confirmCh == nil {
		return nil, fmt.Errorf("no command is pending confirmation")
	}

	s.confirmCh <- req.Msg.Allow
	return connect.NewResponse(&v1.ConfirmShellResponse{}), nil
}
