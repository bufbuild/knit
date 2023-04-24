import { useState } from 'react'
import './App.css'

import type { FilmService } from '../gen/buf/starwars/film/v1/film_knit';
import type { StarshipService } from '../gen/buf/starwars/starship/v1/starship_knit';
import type {  } from '../gen/buf/starwars/relation/v1/relation_knit';
import { createClient } from '@bufbuild/knit';

type Schema = FilmService & StarshipService;

const client = createClient<Schema>({
  baseUrl: 'http://localhost:5173/knit',
})

function App(): JSX.Element {
  const [data, setFilms] = useState(new Array<{title: string, models: string[]}>());

  async function callKnit() {
    const resp = await client.do({
      "buf.starwars.film.v1.FilmService": {
        getFilms: {
          $: { filmIds: ["1","2"] },
          films: {
            title: {},
            starships: {
              $: {},
              model: {},
            },
          },
        },
      },
    });

    const films = resp['buf.starwars.film.v1.FilmService'].getFilms.films.map(film => {
      return {
        title: film.title, 
        models: film.starships.map(ship => ship.model),
      }
    });
    setFilms(films);
  }

  return (
    <div className="App">
      <div>
        <a href="https://github.com/bufbuild/knit" target="_blank">
          <h1>🧶</h1>
        </a>
      </div>
      <h1>Vite + React + Knit</h1>
      <div className="card">
        <button onClick={callKnit}>
          Get Films & Starships
        </button>
      </div>
      <div>
        {data.map((v) => <div key={v.title}>{Film(v.title, v.models)}</div>)}
      </div>
    </div>
  )
}

function Film(title: string, starships: string[]): JSX.Element {
  return (
    <div key={title} className="film">
      <div className="title">{title}</div>
      {starships.map(model => <div key={model}>{model}</div>)}
    </div>
  )
}

export default App
