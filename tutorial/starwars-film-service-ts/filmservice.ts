import { ConnectRouter } from "@bufbuild/connect";
import { FilmService } from "./gen/buf/starwars/film/v1/film_connect";
import { Film, FilmsRequest, FilmsResponse } from "./gen/buf/starwars/film/v1/film_pb";
import { findFilm } from "../starwars-data-ts/film/film";

function routes(router: ConnectRouter) {
    router.service(FilmService, {
        async getFilms(req: FilmsRequest): Promise<FilmsResponse> {
            const films = new Array<Film>();
            req.filmIds.forEach(id => {
                const film = findFilm(id);
                if (film) {
                    films.push(new Film({}));
                }
            })
            return new FilmsResponse({
                films
            })
        }
    })
}