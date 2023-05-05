package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/bufbuild/connect-go"
	grpcreflect "github.com/bufbuild/connect-grpcreflect-go"
	quotev1 "github.com/bufbuild/knit/tutorial/starwars-quote-service-go/gen/buf/starwars/quote/v1"
	"github.com/bufbuild/knit/tutorial/starwars-quote-service-go/gen/buf/starwars/quote/v1/quotev1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	const addr = "127.0.0.1:18003"

	log.Printf("Quote service starting")

	mux := http.NewServeMux()

	path, handler := quotev1connect.NewQuoteServiceHandler(&QuoteService{})
	mux.Handle(path, handler)
	log.Printf("Handling connect service at prefix: %v", path)

	reflector := grpcreflect.NewStaticReflector(
		quotev1connect.QuoteServiceName,
	)

	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

	log.Printf("Listening on: %v", addr)
	err := http.ListenAndServe(
		addr,
		h2c.NewHandler(mux, &http2.Server{}),
	)

	if err != http.ErrServerClosed {
		log.Printf("Error running or stopping: %v", err)
	}
}

type QuoteService struct{}

var _ quotev1connect.QuoteServiceHandler = (*QuoteService)(nil)

func (s *QuoteService) StreamQuotes(
	ctx context.Context,
	req *connect.Request[quotev1.StreamQuotesRequest],
	stream *connect.ServerStream[quotev1.StreamQuotesResponse],
) error {
	log.Println("Starting streaming response...")
	cnt := 0
	for {
		time.Sleep(2 * time.Second)
		resp := &quotev1.StreamQuotesResponse{
			Quote: &quotev1.Quote{
				QuoteId: fmt.Sprintf("%d", cnt),
			},
		}
		err := stream.Send(resp)
		if err != nil {
			log.Printf("Streaming response failed: %v", err)
			return err
		}
	}
}
