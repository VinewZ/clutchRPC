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

	Mu        *sync.Mutex
	ConfirmCh chan bool
}

// UseShell sends a confirmation event to the frontend and waits (with timeout) for the user's answer.
// It then runs `sh -c <command>` and returns its combined output or an error if the command failed/cancelled.
func (s *UseShellServiceServer) UseShell(
	ctx context.Context,
	req *connect.Request[v1.UseShellRequest],
) (*connect.Response[v1.UseShellResponse], error) {
	cmd := req.Msg.Command

	// 1) Quick busy-check under the lock.
	s.Mu.Lock()
	if len(s.ConfirmCh) > 0 {
		s.Mu.Unlock()
		return nil, fmt.Errorf("another confirmation is pending")
	}
	s.Mu.Unlock()

	// 2) Fire the “please confirm” event
	s.App.EmitEvent("clutch:require-confirmation", map[string]string{
		"appName": req.Msg.AppName,
		"command": cmd,
	})
	fmt.Printf("Waiting for confirmation to run command: %q\n", cmd)

	// 3) Wait for either confirmation or timeout
	select {
	case confirmed := <-s.ConfirmCh:
		// reading from the channel automatically empties the buffer
		if !confirmed {
			return nil, fmt.Errorf("command %q cancelled by user", cmd)
		}

	case <-time.After(30 * time.Second):
		// if we time out, drain any late-arriving confirmation
		s.Mu.Lock()
		select {
		case <-s.ConfirmCh:
		default:
		}
		s.Mu.Unlock()

		return nil, fmt.Errorf("confirmation for %q timed out", cmd)
	}

	// 4) Run the shell command
	outBytes, err := exec.Command("sh", "-c", cmd).CombinedOutput()
	output := string(outBytes)
	if err != nil {
		return nil, fmt.Errorf("command %q failed: %w\noutput:\n%s", cmd, err, output)
	}

	return connect.NewResponse(&v1.UseShellResponse{
		Output: output,
	}), nil
}
