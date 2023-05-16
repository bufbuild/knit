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

import { createPromiseClient, ConnectRouter } from "@bufbuild/connect";
import { createConnectTransport } from "@bufbuild/connect-web";
import { StarshipService } from "./gen/buf/starwars/starship/v1/starship_connect";
import { StarshipsRequest } from "./gen/buf/starwars/starship/v1/starship_pb";
import { RelationService } from "./gen/buf/starwars/relation/v1/relation_connect";
import { GetFilmStarshipsRequest, GetFilmStarshipsResponse, GetFilmStarshipsResponse_Result } from "./gen/buf/starwars/relation/v1/relation_pb";


import { fastify } from "fastify";
import { fastifyConnectPlugin } from "@bufbuild/connect-fastify";

const starshipClient = createPromiseClient(StarshipService, createConnectTransport({
    baseUrl: "http://127.0.0.1:18002", // URL of starship service
}));

function routes(router: ConnectRouter) {
    router.service(RelationService, {
        async getFilmStarships(req: GetFilmStarshipsRequest): Promise<GetFilmStarshipsResponse> {
            const values = Array<GetFilmStarshipsResponse_Result>();
            const promises = req.bases.map(async film => {
                const resp = await starshipClient.getStarships(new StarshipsRequest({
                    starshipIds: film.starshipIds,
                }))
                values.push(new GetFilmStarshipsResponse_Result({
                    starships: resp.starships,
                }));
            });
            await Promise.allSettled(promises);
            return new GetFilmStarshipsResponse({ values });
        }
    })
}

async function main() {
    console.log("Relation service starting");

    const server = fastify();
    await server.register(fastifyConnectPlugin, { routes })

    const host = "127.0.0.1";
    const port = 18000;

    console.log(`Listening on: ${host}:${port}`);
    await server.listen({ host, port });
}

void main();