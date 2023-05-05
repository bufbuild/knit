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
	"github.com/bufbuild/knit/tutorial/starwars-data-go/pkg/film"
	filmv1 "github.com/bufbuild/knit/tutorial/starwars-film-service-go/gen/buf/starwars/film/v1"
	"github.com/bufbuild/knit/tutorial/starwars-film-service-go/gen/buf/starwars/film/v1/filmv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	const addr = "127.0.0.1:18001"

	log.Printf("Film service starting")

	mux := http.NewServeMux()

	path, handler := filmv1connect.NewFilmServiceHandler(&FilmService{})
	mux.Handle(path, handler)
	log.Printf("Handling connect service at prefix: %v", path)

	reflector := grpcreflect.NewStaticReflector(
		filmv1connect.FilmServiceName,
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

type FilmService struct{}

var _ filmv1connect.FilmServiceHandler = (*FilmService)(nil)

func (s *FilmService) GetFilms(
	ctx context.Context,
	req *connect.Request[filmv1.FilmsRequest],
) (
	*connect.Response[filmv1.FilmsResponse],
	error,
) {

	resp := &filmv1.FilmsResponse{}
	for _, id := range req.Msg.FilmIds {
		v := film.Find(id)
		if v == nil {
			continue
		}
		resp.Films = append(resp.Films, &filmv1.Film{
			FilmId:        v.FilmID,
			Title:         v.Title,
			EpisodeNumber: v.EpisodeNumber,
			OpeningText:   v.OpeningText,
			Directors:     v.Directors,
			Producers:     v.Producers,
			// Relations
			CharacterIds: v.CharacterIDs,
			StarshipIds:  v.StarshipIDs,
			PlanetIds:    v.PlanetIDs,
			VehicleIds:   v.VehicleIDs,
			SpeciesIds:   v.SpeciesIDs,
		})
	}

	return connect.NewResponse(resp), nil
}
