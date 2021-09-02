module github.com/InVisionApp/opentelemetry-go/exporters/trace/zipkin

go 1.13

replace github.com/InVisionApp/opentelemetry-go => ../../..

require (
	github.com/openzipkin/zipkin-go v0.2.2
	github.com/stretchr/testify v1.4.0
	github.com/InVisionApp/opentelemetry-go v0.4.3
	google.golang.org/grpc v1.27.1
)
