module github.com/InVisionApp/opentelemetry-go/exporters/trace/jaeger

go 1.13

replace github.com/InVisionApp/opentelemetry-go => ../../..

require (
	github.com/apache/thrift v0.13.0
	github.com/google/go-cmp v0.4.0
	github.com/stretchr/testify v1.4.0
	github.com/InVisionApp/opentelemetry-go v0.4.3
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e // indirect
	google.golang.org/api v0.20.0
	google.golang.org/grpc v1.27.1
)
