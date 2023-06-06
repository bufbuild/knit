# Star Wars Film Service in TypeScript

[Back to top of Tutorial]

The film service is not strictly 🧶 Knit, it's just a simple gRPC service
written using [connect-es], but it's needed for the other parts of the
tutorial. Feel free to just run this service and go on to the Knit
specific parts. The film service is made to listen on address
`http://localhost:18001`. Look at the process diagram below
to see where the film service fits into the bigger
picture.

```mermaid
%%{ init: { 'flowchart': { 'curve': 'basis' } } }%%
flowchart LR
A[Knit Client] --> B[Knit Gateway]
subgraph r [Knit Relation Service]
    R{{Relation RPCs}}
end
subgraph f [Film Service]
    F{{Film RPCs}}
end
subgraph s [Starship Service]
    S{{Starship RPCs}}
end
B --> R
B --> F
B --> S
style f stroke:#000,stroke-width:3px
style F stroke:#000,stroke-width:3px
```

## How to run the code
To run the film service clone the repo using `git clone https://github.com/bufbuild/knit.git`,
then execute the following from the base of the repository (the other services must be running too).

[![Slack](https://img.shields.io/badge/If_you_need_help_talk_to_us_in_Slack-Buf-%23e01563)][badges_slack]
```
cd tutorial/starwars-film-service-ts/

npm install
npx tsx fileservice.ts

# Output
Film service starting
Listening on: 127.0.0.1:18001
```

Open another terminal, and at the base directory do:
```
buf curl \
"http://localhost:18001/buf.starwars.film.v1.FilmService/GetFilms" \
--data '{"filmIds":["1"]}' \
--http2-prior-knowledge \
--schema tutorial/starwars-film-service-ts/proto/buf/starwars/film/v1/film.proto
```

[Back to top of Tutorial]: /tutorial
[github.com/bufbuild/knit]: https://github.com/bufbuild/knit
[connect-es]: https://github.com/bufbuild/connect-es
[badges_slack]: https://buf.build/links/slack