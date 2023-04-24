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
	"github.com/bufbuild/knit/tutorial/starwars-data-go/pkg/starship"
	starshipv1 "github.com/bufbuild/knit/tutorial/starwars-starship-service-go/gen/buf/starwars/starship/v1"
	"github.com/bufbuild/knit/tutorial/starwars-starship-service-go/gen/buf/starwars/starship/v1/starshipv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	const addr = "127.0.0.1:18002"

	log.Printf("Starship service starting")

	mux := http.NewServeMux()

	path, handler := starshipv1connect.NewStarshipServiceHandler(&StarshipService{})
	mux.Handle(path, handler)
	log.Printf("Handling connect service at prefix: %v", path)

	reflector := grpcreflect.NewStaticReflector(
		starshipv1connect.StarshipServiceName,
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

type StarshipService struct{}

var _ starshipv1connect.StarshipServiceHandler = (*StarshipService)(nil)

func (s *StarshipService) GetStarships(
	ctx context.Context,
	req *connect.Request[starshipv1.StarshipsRequest],
) (
	*connect.Response[starshipv1.StarshipsResponse],
	error,
) {

	resp := &starshipv1.StarshipsResponse{}
	for _, id := range req.Msg.StarshipIds {
		v := starship.Find(id)
		if v == nil {
			continue
		}
		resp.Starships = append(resp.Starships, &starshipv1.Starship{
			StarshipId:           v.StarshipID,
			Name:                 v.Name,
			Model:                v.Model,
			Manufacturer:         v.Manufacturer,
			CostInCredits:        v.CostInCredits,
			Length:               v.Length,
			MaxAtmospheringSpeed: v.MaxAtmospheringSpeed,
			Crew:                 v.Crew,
			Passengers:           v.Passengers,
			CargoCapacity:        v.CargoCapacity,
			Consumables:          v.Consumables,
			HyperdriveRating:     float32(v.HyperdriveRating),
			Mglt:                 v.MGLT,
			StarshipClass:        v.StarshipClass,
			// Relations
			PilotIds: v.PilotIDs,
			FilmIds:  v.FilmIDs,
		})
	}

	return connect.NewResponse(resp), nil
}
