module github.com/InVisionApp/opentelemetry-go/example/jaeger

go 1.13

replace (
	github.com/InVisionApp/opentelemetry-go => ../..
	github.com/InVisionApp/opentelemetry-go/exporters/trace/jaeger => ../../exporters/trace/jaeger
)

require (
	github.com/InVisionApp/opentelemetry-go v0.4.3
	github.com/InVisionApp/opentelemetry-go/exporters/trace/jaeger v0.4.3
)
