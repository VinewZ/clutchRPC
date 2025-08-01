// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: clutch/v1/toggle_window.proto

package v1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/vinewz/clutchRPC/go/gen/clutch/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// ToggleWindowServiceName is the fully-qualified name of the ToggleWindowService service.
	ToggleWindowServiceName = "clutch_rpc.v1.ToggleWindowService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ToggleWindowServiceToggleWindowProcedure is the fully-qualified name of the ToggleWindowService's
	// ToggleWindow RPC.
	ToggleWindowServiceToggleWindowProcedure = "/clutch_rpc.v1.ToggleWindowService/ToggleWindow"
)

// ToggleWindowServiceClient is a client for the clutch_rpc.v1.ToggleWindowService service.
type ToggleWindowServiceClient interface {
	ToggleWindow(context.Context, *connect.Request[v1.ToggleWindowRequest]) (*connect.Response[v1.ToggleWindowResponse], error)
}

// NewToggleWindowServiceClient constructs a client for the clutch_rpc.v1.ToggleWindowService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewToggleWindowServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ToggleWindowServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	toggleWindowServiceMethods := v1.File_clutch_v1_toggle_window_proto.Services().ByName("ToggleWindowService").Methods()
	return &toggleWindowServiceClient{
		toggleWindow: connect.NewClient[v1.ToggleWindowRequest, v1.ToggleWindowResponse](
			httpClient,
			baseURL+ToggleWindowServiceToggleWindowProcedure,
			connect.WithSchema(toggleWindowServiceMethods.ByName("ToggleWindow")),
			connect.WithClientOptions(opts...),
		),
	}
}

// toggleWindowServiceClient implements ToggleWindowServiceClient.
type toggleWindowServiceClient struct {
	toggleWindow *connect.Client[v1.ToggleWindowRequest, v1.ToggleWindowResponse]
}

// ToggleWindow calls clutch_rpc.v1.ToggleWindowService.ToggleWindow.
func (c *toggleWindowServiceClient) ToggleWindow(ctx context.Context, req *connect.Request[v1.ToggleWindowRequest]) (*connect.Response[v1.ToggleWindowResponse], error) {
	return c.toggleWindow.CallUnary(ctx, req)
}

// ToggleWindowServiceHandler is an implementation of the clutch_rpc.v1.ToggleWindowService service.
type ToggleWindowServiceHandler interface {
	ToggleWindow(context.Context, *connect.Request[v1.ToggleWindowRequest]) (*connect.Response[v1.ToggleWindowResponse], error)
}

// NewToggleWindowServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewToggleWindowServiceHandler(svc ToggleWindowServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	toggleWindowServiceMethods := v1.File_clutch_v1_toggle_window_proto.Services().ByName("ToggleWindowService").Methods()
	toggleWindowServiceToggleWindowHandler := connect.NewUnaryHandler(
		ToggleWindowServiceToggleWindowProcedure,
		svc.ToggleWindow,
		connect.WithSchema(toggleWindowServiceMethods.ByName("ToggleWindow")),
		connect.WithHandlerOptions(opts...),
	)
	return "/clutch_rpc.v1.ToggleWindowService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ToggleWindowServiceToggleWindowProcedure:
			toggleWindowServiceToggleWindowHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedToggleWindowServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedToggleWindowServiceHandler struct{}

func (UnimplementedToggleWindowServiceHandler) ToggleWindow(context.Context, *connect.Request[v1.ToggleWindowRequest]) (*connect.Response[v1.ToggleWindowResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("clutch_rpc.v1.ToggleWindowService.ToggleWindow is not implemented"))
}
