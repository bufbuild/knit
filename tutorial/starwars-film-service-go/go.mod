module github.com/bufbuild/knit/tutorial/starwars-film-service-go

go 1.20

replace github.com/bufbuild/knit/tutorial/starwars-data-go => ../starwars-data-go

require (
	github.com/bufbuild/connect-go v1.0.0
	github.com/bufbuild/connect-grpcreflect-go v1.0.0
	github.com/bufbuild/knit/tutorial/starwars-data-go v0.0.0-20230505135109-8bf7f257e49d
	golang.org/x/net v0.9.0
	google.golang.org/protobuf v1.33.0
)

require golang.org/x/text v0.9.0 // indirect
