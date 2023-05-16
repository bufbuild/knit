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