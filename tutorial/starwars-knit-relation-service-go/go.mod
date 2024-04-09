module github.com/bufbuild/knit/tutorial/starwars-knit-relation-service-go

go 1.20

require (
	buf.build/gen/go/bufbuild/knit/protocolbuffers/go v1.30.0-20230504140941-3dc602456973.1
	github.com/bufbuild/connect-go v1.7.0
	github.com/bufbuild/connect-grpcreflect-go v1.0.1-0.20230317120624-8e24e9604364
	golang.org/x/net v0.9.0
	google.golang.org/protobuf v1.33.0
)

require golang.org/x/text v0.9.0 // indirect
