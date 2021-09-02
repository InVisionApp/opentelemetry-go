module go.opentelemetry.go/otel/example/zipkin

go 1.13

replace (
	github.com/InVisionApp/opentelemetry-go => ../..
	github.com/InVisionApp/opentelemetry-go/exporters/trace/zipkin => ../../exporters/trace/zipkin
)

require (
	github.com/InVisionApp/opentelemetry-go v0.4.3
	github.com/InVisionApp/opentelemetry-go/exporters/trace/zipkin v0.4.3
)
