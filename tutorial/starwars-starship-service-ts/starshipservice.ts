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

import { ConnectRouter } from "@bufbuild/connect";
import { StarshipService } from "./gen/buf/starwars/starship/v1/starship_connect";
import { Starship, StarshipsRequest, StarshipsResponse } from "./gen/buf/starwars/starship/v1/starship_pb";
import { findStarship } from "../starwars-data-ts/starship/starship";

import { fastify } from "fastify";
import { fastifyConnectPlugin } from "@bufbuild/connect-fastify";

function routes(router: ConnectRouter) {
    router.service(StarshipService, {
        async getStarships(req: StarshipsRequest): Promise<StarshipsResponse> {
            console.log(`Getting starships: ${req.starshipIds}`)
            const starships = new Array<Starship>();
            req.starshipIds.forEach(id => {
                const starship = findStarship(id);
                if (starship) {
                    starships.push(new Starship({
                        starshipId: starship.starshipId,
                        name: starship.name,
                        model: starship.model,
                        manufacturer: starship.manufacturer,
                        costInCredits: BigInt(starship.costInCredits),
                        length: BigInt(starship.length),
                        maxAtmospheringSpeed: BigInt(starship.maxAtmospheringSpeed),
                        crew: BigInt(starship.crew),
                        passengers: BigInt(starship.passengers),
                        cargoCapacity: BigInt(starship.cargoCapacity),
                        consumables: starship.consumables,
                        hyperdriveRating: starship.hyperdriveRating,
                        mglt: BigInt(starship.mglt),
                        starshipClass: starship.starshipClass,
                        pilotIds: starship.pilotIds,
                        filmIds: starship.filmIds,
                    }));
                }
            })
            return new StarshipsResponse({ starships });
        }
    });
}

async function main() {
    console.log("Starship service starting");

    const server = fastify();
    await server.register(fastifyConnectPlugin, { routes })

    const host = "127.0.0.1";
    const port = 18002;

    console.log(`Listening on: ${host}:${port}`);
    await server.listen({ host, port });
}

void main();