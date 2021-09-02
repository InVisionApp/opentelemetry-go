// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package simple_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/InVisionApp/opentelemetry-go/api/core"
	"github.com/InVisionApp/opentelemetry-go/api/metric"
	"github.com/InVisionApp/opentelemetry-go/sdk/metric/aggregator/array"
	"github.com/InVisionApp/opentelemetry-go/sdk/metric/aggregator/ddsketch"
	"github.com/InVisionApp/opentelemetry-go/sdk/metric/aggregator/histogram"
	"github.com/InVisionApp/opentelemetry-go/sdk/metric/aggregator/minmaxsumcount"
	"github.com/InVisionApp/opentelemetry-go/sdk/metric/aggregator/sum"
	"github.com/InVisionApp/opentelemetry-go/sdk/metric/selector/simple"
)

var (
	testCounterDesc  = metric.NewDescriptor("counter", metric.CounterKind, core.Int64NumberKind)
	testMeasureDesc  = metric.NewDescriptor("measure", metric.MeasureKind, core.Int64NumberKind)
	testObserverDesc = metric.NewDescriptor("observer", metric.ObserverKind, core.Int64NumberKind)
)

func TestInexpensiveMeasure(t *testing.T) {
	inex := simple.NewWithInexpensiveMeasure()
	require.NotPanics(t, func() { _ = inex.AggregatorFor(&testCounterDesc).(*sum.Aggregator) })
	require.NotPanics(t, func() { _ = inex.AggregatorFor(&testMeasureDesc).(*minmaxsumcount.Aggregator) })
	require.NotPanics(t, func() { _ = inex.AggregatorFor(&testObserverDesc).(*minmaxsumcount.Aggregator) })
}

func TestSketchMeasure(t *testing.T) {
	sk := simple.NewWithSketchMeasure(ddsketch.NewDefaultConfig())
	require.NotPanics(t, func() { _ = sk.AggregatorFor(&testCounterDesc).(*sum.Aggregator) })
	require.NotPanics(t, func() { _ = sk.AggregatorFor(&testMeasureDesc).(*ddsketch.Aggregator) })
	require.NotPanics(t, func() { _ = sk.AggregatorFor(&testObserverDesc).(*ddsketch.Aggregator) })
}

func TestExactMeasure(t *testing.T) {
	ex := simple.NewWithExactMeasure()
	require.NotPanics(t, func() { _ = ex.AggregatorFor(&testCounterDesc).(*sum.Aggregator) })
	require.NotPanics(t, func() { _ = ex.AggregatorFor(&testMeasureDesc).(*array.Aggregator) })
	require.NotPanics(t, func() { _ = ex.AggregatorFor(&testObserverDesc).(*array.Aggregator) })
}

func TestHistogramMeasure(t *testing.T) {
	ex := simple.NewWithHistogramMeasure([]core.Number{})
	require.NotPanics(t, func() { _ = ex.AggregatorFor(&testCounterDesc).(*sum.Aggregator) })
	require.NotPanics(t, func() { _ = ex.AggregatorFor(&testMeasureDesc).(*histogram.Aggregator) })
	require.NotPanics(t, func() { _ = ex.AggregatorFor(&testObserverDesc).(*histogram.Aggregator) })
}
