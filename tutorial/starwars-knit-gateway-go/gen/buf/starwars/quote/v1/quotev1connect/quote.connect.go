// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: buf/starwars/quote/v1/quote.proto

package quotev1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/bufbuild/knit/tutorial/starwars-knit-gateway-go/gen/buf/starwars/quote/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// QuoteServiceName is the fully-qualified name of the QuoteService service.
	QuoteServiceName = "buf.starwars.quote.v1.QuoteService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// QuoteServiceStreamQuotesProcedure is the fully-qualified name of the QuoteService's StreamQuotes
	// RPC.
	QuoteServiceStreamQuotesProcedure = "/buf.starwars.quote.v1.QuoteService/StreamQuotes"
)

// QuoteServiceClient is a client for the buf.starwars.quote.v1.QuoteService service.
type QuoteServiceClient interface {
	StreamQuotes(context.Context, *connect_go.Request[v1.StreamQuotesRequest]) (*connect_go.ServerStreamForClient[v1.StreamQuotesResponse], error)
}

// NewQuoteServiceClient constructs a client for the buf.starwars.quote.v1.QuoteService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewQuoteServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) QuoteServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &quoteServiceClient{
		streamQuotes: connect_go.NewClient[v1.StreamQuotesRequest, v1.StreamQuotesResponse](
			httpClient,
			baseURL+QuoteServiceStreamQuotesProcedure,
			opts...,
		),
	}
}

// quoteServiceClient implements QuoteServiceClient.
type quoteServiceClient struct {
	streamQuotes *connect_go.Client[v1.StreamQuotesRequest, v1.StreamQuotesResponse]
}

// StreamQuotes calls buf.starwars.quote.v1.QuoteService.StreamQuotes.
func (c *quoteServiceClient) StreamQuotes(ctx context.Context, req *connect_go.Request[v1.StreamQuotesRequest]) (*connect_go.ServerStreamForClient[v1.StreamQuotesResponse], error) {
	return c.streamQuotes.CallServerStream(ctx, req)
}

// QuoteServiceHandler is an implementation of the buf.starwars.quote.v1.QuoteService service.
type QuoteServiceHandler interface {
	StreamQuotes(context.Context, *connect_go.Request[v1.StreamQuotesRequest], *connect_go.ServerStream[v1.StreamQuotesResponse]) error
}

// NewQuoteServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewQuoteServiceHandler(svc QuoteServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle(QuoteServiceStreamQuotesProcedure, connect_go.NewServerStreamHandler(
		QuoteServiceStreamQuotesProcedure,
		svc.StreamQuotes,
		opts...,
	))
	return "/buf.starwars.quote.v1.QuoteService/", mux
}

// UnimplementedQuoteServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedQuoteServiceHandler struct{}

func (UnimplementedQuoteServiceHandler) StreamQuotes(context.Context, *connect_go.Request[v1.StreamQuotesRequest], *connect_go.ServerStream[v1.StreamQuotesResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("buf.starwars.quote.v1.QuoteService.StreamQuotes is not implemented"))
}