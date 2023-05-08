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
	"context"
	"log"
	"net/http"

	"github.com/bufbuild/connect-go"
	grpcreflect "github.com/bufbuild/connect-grpcreflect-go"
	filmv1 "github.com/bufbuild/knit/tutorial/starwars-knit-relation-service-go/gen/buf/starwars/film/v1"
	"github.com/bufbuild/knit/tutorial/starwars-knit-relation-service-go/gen/buf/starwars/film/v1/filmv1connect"
	relationv1 "github.com/bufbuild/knit/tutorial/starwars-knit-relation-service-go/gen/buf/starwars/relation/v1"
	"github.com/bufbuild/knit/tutorial/starwars-knit-relation-service-go/gen/buf/starwars/relation/v1/relationv1connect"
	starshipv1 "github.com/bufbuild/knit/tutorial/starwars-knit-relation-service-go/gen/buf/starwars/starship/v1"
	"github.com/bufbuild/knit/tutorial/starwars-knit-relation-service-go/gen/buf/starwars/starship/v1/starshipv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	const addr = "127.0.0.1:18000"

	log.Printf("Knit relation service starting")

	rel := &RelationService{
		filmClient:     filmv1connect.NewFilmServiceClient(http.DefaultClient, "http://localhost:18001"),
		starshipClient: starshipv1connect.NewStarshipServiceClient(http.DefaultClient, "http://localhost:18002"),
	}

	path, handler := relationv1connect.NewRelationServiceHandler(rel)
	mux := http.NewServeMux()
	mux.Handle(path, handler)
	log.Printf("Handling connect service at prefix: %v", path)

	reflector := grpcreflect.NewStaticReflector(
		relationv1connect.RelationServiceName,
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

type RelationService struct {
	filmClient     filmv1connect.FilmServiceClient
	starshipClient starshipv1connect.StarshipServiceClient
}

var _ (relationv1connect.RelationServiceHandler) = (*RelationService)(nil)

func (s *RelationService) GetFilmStarships(
	ctx context.Context,
	req *connect.Request[relationv1.GetFilmStarshipsRequest],
) (
	*connect.Response[relationv1.GetFilmStarshipsResponse],
	error,
) {
	resp := relationv1.GetFilmStarshipsResponse{}

	for _, film := range req.Msg.Bases {
		starships, err := s.starshipClient.GetStarships(ctx, connect.NewRequest(&starshipv1.StarshipsRequest{
			StarshipIds: film.StarshipIds,
		}))
		if err != nil {
			log.Printf("Request to starship service with %v failed: %v", film.StarshipIds, err)
			return nil, err
		}

		resp.Values = append(resp.Values, &relationv1.GetFilmStarshipsResponse_Value{
			Starships: starships.Msg.Starships,
		})
	}

	return connect.NewResponse(&resp), nil
}

func (s *RelationService) GetQuoteFilm(
	ctx context.Context,
	req *connect.Request[relationv1.GetQuoteFilmRequest],
) (
	*connect.Response[relationv1.GetQuoteFilmResponse],
	error,
) {
	resp := relationv1.GetQuoteFilmResponse{}

	for _, quote := range req.Msg.Bases {
		films, err := s.filmClient.GetFilms(ctx, connect.NewRequest(&filmv1.FilmsRequest{
			FilmIds: []string{quote.FilmId},
		}))
		if err != nil {
			log.Printf("Request to film service with %v failed: %v", quote.FilmId, err)
			return nil, err
		}

		if len(films.Msg.Films) == 0 {
			log.Printf("Request to film service with %v returned zero films", quote.FilmId)
			continue
		}
		resp.Values = append(resp.Values, &relationv1.GetQuoteFilmResponse_Value{
			Film: films.Msg.Films[0],
		})
	}

	return connect.NewResponse(&resp), nil
}
