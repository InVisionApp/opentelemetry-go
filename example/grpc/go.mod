module github.com/InVisionApp/opentelemetry-go/example/grpc

go 1.13

replace github.com/InVisionApp/opentelemetry-go => ../..

require (
	github.com/golang/protobuf v1.3.2
	github.com/InVisionApp/opentelemetry-go v0.4.3
	golang.org/x/net v0.0.0-20190311183353-d8887717615a
	google.golang.org/grpc v1.27.1
)
