import { useEffect, useState } from 'react'
import './App.css'

import type { FilmService } from '../gen/buf/starwars/film/v1/film_knit';
import type { StarshipService } from '../gen/buf/starwars/starship/v1/starship_knit';
import type { Quote, QuoteService } from '../gen/buf/starwars/quote/v1/quote_knit';
import type {  } from '../gen/buf/starwars/relation/v1/relation_knit';
import { createClient, Mask, Query } from '@bufbuild/knit';

type Schema = FilmService & StarshipService & QuoteService;

const client = createClient<Schema>({
  baseUrl: 'http://localhost:5173/knit',
})

function App(): JSX.Element {
  return (
    <div className="App">
      <ListenComponent />
    </div>
  )
}

function FetchComponent(): JSX.Element {
  const [data, setFilms] = useState(new Array<{title: string, models: string[]}>());

  async function getFilmStarships() {
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
    <div></div>
  )
}

function ListenComponent(): JSX.Element {
  return (
    <div>
      <div className="instruction">
        Knit can stream results from the backend using the client's <code>listen()</code> method,
        which calls through the Knit gateway to any streaming RPC like <code>StreamQuotes()</code>.
        If a relation exists in the query, like the field <code>film</code> here, the RPCs needed
        to fetch the related data are called by the Knit gateway each time a new response is emitted
        on the stream. 
      </div>
      <div className="twopanel">
        <div className="panel" style={{width:"45%"}}>
          <pre><code>
            {streamingQueryExample}
            </code></pre>
        </div>
        <div className="panel" style={{width:"55%"}}>
          <pre>
            {streamingRpcExample}
          </pre>
        </div>
      </div>
      <div className="controls">
          <QuoteStreamComponent />
      </div>
    </div>
  )
}

const quoteQuery = {
  text: {},
  film: {
    title: {},
  }
} satisfies Query<Quote>;

type QueryResult = Mask<typeof quoteQuery, Quote>

function QuoteStreamComponent(): JSX.Element {
  const [quote, setQuote] = useState<QueryResult | null>(null);
  const [allowStreaming, setAllowStreaming] = useState(false);

  useEffect(() => {
    if (!allowStreaming) return;
    let stopStream = false;
    const stream = async () => {
        const quoteStream = client.listen({
        "buf.starwars.quote.v1.QuoteService": {
          streamQuotes: {
            $: { limit: 5n },
            quote: {              
              ...quoteQuery,
            },
          },
        }
      });
      for await (const resp of quoteStream) {
        if (stopStream) {
          return;
        }
        const quote = resp['buf.starwars.quote.v1.QuoteService'].streamQuotes.quote;        
        setQuote(quote!);
      }
    }
    stream();
    return () => { stopStream = true }
  }, [allowStreaming])

  return (
    <div>
      <div>
        <button onClick={() => setAllowStreaming(true)}>
          Start quotes
        </button>
        <button onClick={() => setAllowStreaming(false)}>
          Stop quotes
        </button>
      </div>
      <div>
        <QuoteComponent quote={quote?.text} film={quote?.film?.title}/>
      </div>
    </div>
  )
}

function QuoteComponent({quote, film}: {quote: string | undefined, film: string | undefined}): JSX.Element {
  if (quote && film) {
    console.log("Quote and film: ", quote, film);
    return (
      <div>
        <span className="quotetext">“{quote}”</span>, from {film}
      </div>
    )
  } else {
    return (
      <div></div>
    )
  }
}

export default App

const streamingQueryExample = `
async function streamQuotes() {
  const quoteStream = client.listen({
    "buf.starwars.quote.v1.QuoteService": {
      streamQuotes: {
        $: { limit: 5 },
        quote: {
          text: {},
          film: {
            $: {},
            title: {},
          }
        },
      },
    },
  });

  for await (const quote of quoteStream) {
    console.log(quote.text, quote.film.title);
  }
}
`

const streamingRpcExample = `
message Quote {
  string quote_id = 1;
  string text = 2;
  string person_id = 3;
  string film_id = 4;
}

message StreamQuotesRequest {}

message StreamQuotesResponse {
  Quote quote = 1;
}

service QuoteService {
  rpc StreamQuotes(StreamQuotesRequest) returns (stream StreamQuotesResponse) {}
}
`