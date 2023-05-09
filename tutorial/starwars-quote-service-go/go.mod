module github.com/bufbuild/knit/tutorial/starwars-quote-service-go

go 1.20

replace github.com/bufbuild/knit/tutorial/starwars-data-go => ../starwars-data-go

require (
	github.com/bufbuild/connect-go v1.7.0
	github.com/bufbuild/connect-grpcreflect-go v1.0.0
	github.com/bufbuild/knit/tutorial/starwars-data-go v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.9.0
	google.golang.org/protobuf v1.30.0
)

require golang.org/x/text v0.9.0 // indirect
