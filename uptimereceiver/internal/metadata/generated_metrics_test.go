// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver/receivertest"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

type testDataSet int

const (
	testDataSetDefault testDataSet = iota
	testDataSetAll
	testDataSetNone
)

func TestMetricsBuilder(t *testing.T) {
	tests := []struct {
		name        string
		metricsSet  testDataSet
		resAttrsSet testDataSet
		expectEmpty bool
	}{
		{
			name: "default",
		},
		{
			name:        "all_set",
			metricsSet:  testDataSetAll,
			resAttrsSet: testDataSetAll,
		},
		{
			name:        "none_set",
			metricsSet:  testDataSetNone,
			resAttrsSet: testDataSetNone,
			expectEmpty: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			start := pcommon.Timestamp(1_000_000_000)
			ts := pcommon.Timestamp(1_000_001_000)
			observedZapCore, observedLogs := observer.New(zap.WarnLevel)
			settings := receivertest.NewNopSettings()
			settings.Logger = zap.New(observedZapCore)
			mb := NewMetricsBuilder(loadMetricsBuilderConfig(t, test.name), settings, WithStartTime(start))

			expectedWarnings := 0

			assert.Equal(t, expectedWarnings, observedLogs.Len())

			defaultMetricsCount := 0
			allMetricsCount := 0

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordUptimeDataPoint(ts, 1, "hostname-val", "os-val", "arch-val", "type-val")

			res := pcommon.NewResource()
			metrics := mb.Emit(WithResource(res))

			if test.expectEmpty {
				assert.Equal(t, 0, metrics.ResourceMetrics().Len())
				return
			}

			assert.Equal(t, 1, metrics.ResourceMetrics().Len())
			rm := metrics.ResourceMetrics().At(0)
			assert.Equal(t, res, rm.Resource())
			assert.Equal(t, 1, rm.ScopeMetrics().Len())
			ms := rm.ScopeMetrics().At(0).Metrics()
			if test.metricsSet == testDataSetDefault {
				assert.Equal(t, defaultMetricsCount, ms.Len())
			}
			if test.metricsSet == testDataSetAll {
				assert.Equal(t, allMetricsCount, ms.Len())
			}
			validatedMetrics := make(map[string]bool)
			for i := 0; i < ms.Len(); i++ {
				switch ms.At(i).Name() {
				case "uptime":
					assert.False(t, validatedMetrics["uptime"], "Found a duplicate in the metrics slice: uptime")
					validatedMetrics["uptime"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "System uptime in seconds", ms.At(i).Description())
					assert.Equal(t, "seconds", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("hostname")
					assert.True(t, ok)
					assert.EqualValues(t, "hostname-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("os")
					assert.True(t, ok)
					assert.EqualValues(t, "os-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("arch")
					assert.True(t, ok)
					assert.EqualValues(t, "arch-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("type")
					assert.True(t, ok)
					assert.EqualValues(t, "type-val", attrVal.Str())
				}
			}
		})
	}
}
