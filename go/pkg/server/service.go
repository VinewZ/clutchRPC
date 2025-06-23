package clutchrpc

import (
	"log"
	"net/http"

	"github.com/vinewz/clutchRPC/go/internal/server"
	pbConnect "github.com/vinewz/clutchRPC/go/pb/clutch/v1/v1connect"
	"github.com/wailsapp/wails/v3/pkg/application"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	connectcors "connectrpc.com/cors"
	"github.com/rs/cors"
)

// ClutchServer encapsulates the gRPC service plus confirmation logic.
type ClutchServer struct {
	server.ToggleWindowServiceServer
	server.UseShellServiceServer
	server.SayHiServiceServer
}

func withCORS(h http.Handler) http.Handler {
	return cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: connectcors.AllowedMethods(),
		AllowedHeaders: connectcors.AllowedHeaders(),
		ExposedHeaders: connectcors.ExposedHeaders(),
	}).Handler(h)
}

// New returns a ClutchServer ready to be registered.
// `timeout` controls how long to wait for user confirmation.
func New(app *application.App) *ClutchServer {
	return &ClutchServer{
		ToggleWindowServiceServer: server.ToggleWindowServiceServer{App: app},
		UseShellServiceServer:     server.UseShellServiceServer{App: app},
		SayHiServiceServer:     server.SayHiServiceServer{},
	}
}

func (s *ClutchServer) ListenAndServe(addr string) error {
	mux := http.NewServeMux()
	{
		path, handler := pbConnect.NewSayHiServiceHandler(&s.SayHiServiceServer)
		mux.Handle(path, handler)
	}
	{
		path, handler := pbConnect.NewUseShellServiceHandler(&s.UseShellServiceServer)
		mux.Handle(path, handler)
	}
	{
		path, handler := pbConnect.NewToggleWindowServiceHandler(&s.ToggleWindowServiceServer)
		mux.Handle(path, handler)
	}
	corsMux := withCORS(mux)

	finalHandler := h2c.NewHandler(corsMux, &http2.Server{})

	log.Println("Started on", addr, "with CORS enabled")
	err := http.ListenAndServe(addr, finalHandler)
	if err != nil {
		return err
	}
	return nil
}
