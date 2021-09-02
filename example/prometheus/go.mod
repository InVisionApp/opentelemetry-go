module github.com/InVisionApp/opentelemetry-go/example/prometheus

go 1.13

replace (
	github.com/InVisionApp/opentelemetry-go => ../..
	github.com/InVisionApp/opentelemetry-go/exporters/metric/prometheus => ../../exporters/metric/prometheus
)

require (
	github.com/InVisionApp/opentelemetry-go v0.4.3
	github.com/InVisionApp/opentelemetry-go/exporters/metric/prometheus v0.4.3
)
