package ismetrics

import (
	"time"

	"github.com/cgascoig/intersight-metrics/pkg/intersight"
)

func init() {
	druidGroupByRequest := intersight.TelemetryDruidGroupByRequest{
		QueryType: "groupBy",
		Granularity: intersight.TelemetryDruidGranularity{
			TelemetryDruidPeriodGranularity: &intersight.TelemetryDruidPeriodGranularity{
				Type:   "period",
				Period: "PT15M",
			},
		},
		DataSource: intersight.TelemetryDruidDataSource{
			TelemetryDruidTableDataSource: &intersight.TelemetryDruidTableDataSource{
				Name: "hx",
				Type: "table",
			},
		},
		Dimensions: []intersight.TelemetryDruidDimensionSpec{
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "deviceId",
					OutputName: "deviceId",
					OutputType: "STRING",
				},
			},
		},
		Intervals: []string{getIntervalString(time.Now(), time.Duration(15)*time.Minute)},
		Aggregations: &[]intersight.TelemetryDruidAggregator{
			{
				TelemetryDruidSumAggregator: &intersight.TelemetryDruidSumAggregator{
					Type:      "doubleSum",
					Name:      "sumReadBytes",
					FieldName: "sumReadBytes",
				},
			},
			{
				TelemetryDruidSumAggregator: &intersight.TelemetryDruidSumAggregator{
					Type:      "doubleSum",
					Name:      "sumWriteBytes",
					FieldName: "sumWriteBytes",
				},
			},
		},
	}

	c := NewDruidCollector("hx", druidGroupByRequest, []string{"deviceId"})
	c.RegisterMetric("intersight_hx_read_bytes", "HX bytes read", "sumReadBytes")
	c.RegisterMetric("intersight_hx_write_bytes", "HX bytes written", "sumWriteBytes")
	RegisterCollector(c)
}
