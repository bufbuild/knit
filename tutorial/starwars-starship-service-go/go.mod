module github.com/bufbuild/knit/tutorial/starwars-starship-service-go

go 1.23.0

require (
	github.com/bufbuild/connect-go v1.6.0
	github.com/bufbuild/connect-grpcreflect-go v1.0.0
	github.com/bufbuild/knit/tutorial/starwars-data-go v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.39.0
	google.golang.org/protobuf v1.33.0
)

replace github.com/bufbuild/knit/tutorial/starwars-data-go => ../starwars-data-go

require golang.org/x/text v0.24.0 // indirect
