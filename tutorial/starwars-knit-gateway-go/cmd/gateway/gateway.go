// Copyright 2023 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"log"
	"net/http"
	"net/url"

	grpcreflect "github.com/bufbuild/connect-grpcreflect-go"
	"github.com/bufbuild/knit-go"
	"github.com/bufbuild/knit/tutorial/starwars-knit-gateway-go/gen/buf/starwars/film/v1/filmv1connect"
	"github.com/bufbuild/knit/tutorial/starwars-knit-gateway-go/gen/buf/starwars/quote/v1/quotev1connect"
	"github.com/bufbuild/knit/tutorial/starwars-knit-gateway-go/gen/buf/starwars/relation/v1/relationv1connect"
	"github.com/bufbuild/knit/tutorial/starwars-knit-gateway-go/gen/buf/starwars/starship/v1/starshipv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func main() {
	const addr = "127.0.0.1:8080"

	log.Printf("Knit gateway starting")

	relationServiceURL, err := url.Parse("http://127.0.0.1:18000")
	if err != nil {
		log.Fatalf("Failed to parse relation URL: %v", err)
	}

	filmServiceURL, err := url.Parse("http://127.0.0.1:18001")
	if err != nil {
		log.Fatalf("Failed to parse film URL: %v", err)
	}

	starshipServiceURL, err := url.Parse("http://127.0.0.1:18002")
	if err != nil {
		log.Fatalf("Failed to parse person URL: %v", err)
	}

	quoteServiceURL, err := url.Parse("http://127.0.0.1:18003")
	if err != nil {
		log.Fatalf("Failed to parse person URL: %v", err)
	}

	gateway := knit.Gateway{}
	if err := gateway.AddServiceByName(protoreflect.FullName(relationv1connect.RelationServiceName), knit.WithRoute(relationServiceURL)); err != nil {
		log.Fatalf("Failed to add service: %v, error: %v", relationv1connect.RelationServiceName, err)
	}
	if err := gateway.AddServiceByName(protoreflect.FullName(filmv1connect.FilmServiceName), knit.WithRoute(filmServiceURL)); err != nil {
		log.Fatalf("Failed to add service: %v, error: %v", filmv1connect.FilmServiceName, err)
	}
	if err := gateway.AddServiceByName(protoreflect.FullName(starshipv1connect.StarshipServiceName), knit.WithRoute(starshipServiceURL)); err != nil {
		log.Fatalf("Failed to add service: %v, error: %v", starshipv1connect.StarshipServiceName, err)
	}
	if err := gateway.AddServiceByName(protoreflect.FullName(quotev1connect.QuoteServiceName), knit.WithRoute(quoteServiceURL)); err != nil {
		log.Fatalf("Failed to add service: %v, error: %v", quotev1connect.QuoteServiceName, err)
	}

	mux := http.NewServeMux()
	mux.Handle(gateway.AsHandler())

	reflector := grpcreflect.NewStaticReflector(
		relationv1connect.RelationServiceName,
		filmv1connect.FilmServiceName,
		starshipv1connect.StarshipServiceName,
	)

	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

	log.Printf("Listening on: %v", addr)
	err = http.ListenAndServe(
		addr,
		h2c.NewHandler(mux, &http2.Server{}),
	)

	if err != http.ErrServerClosed {
		log.Printf("Error running or stopping: %v", err)
	}
}
