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
import { FilmService } from "./gen/buf/starwars/film/v1/film_connect";
import { Film, FilmsRequest, FilmsResponse } from "./gen/buf/starwars/film/v1/film_pb";
import { findFilm } from "../starwars-data-ts/film/film";

import { fastify } from "fastify";
import { fastifyConnectPlugin } from "@bufbuild/connect-fastify";

function routes(router: ConnectRouter) {
    router.service(FilmService, {
        async getFilms(req: FilmsRequest): Promise<FilmsResponse> {
            console.log(`Getting films: ${req.filmIds}`)
            const films = new Array<Film>();
            req.filmIds.forEach(id => {
                const film = findFilm(id);
                if (film) {
                    films.push(new Film({
                        filmId: film.filmId,
                        episodeNumber: BigInt(film.episodeNumber),
                        directors: film.directors,
                        producers: film.producers,
                        openingText: film.openingText,
                        title: film.title,
                        characterIds: film.characterIds,
                        planetIds: film.planetIds,
                        speciesIds: film.speciesIds,
                        starshipIds: film.starshipIds,
                        vehicleIds: film.vehicleIds,
                    }));
                }
            })
            return new FilmsResponse({ films })
        }
    })
}

async function main() {
    console.log("Film service starting");

    const server = fastify();
    await server.register(fastifyConnectPlugin, { routes })

    const host = "127.0.0.1";
    const port = 18001;

    console.log(`Listening on: ${host}:${port}`);
    await server.listen({ host, port });
}

void main();