package ismetrics

import (
	"time"

	"github.com/cgascoig/intersight-prometheus-exporter/pkg/intersight"
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
				Name: "ucs_ether_port_stat",
				Type: "table",
			},
		},
		Dimensions: []intersight.TelemetryDruidDimensionSpec{
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "dn",
					OutputName: "port",
					OutputType: "STRING",
				},
			},
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
					Type:      "longSum",
					Name:      "traffic",
					FieldName: "sumBytesTx",
				},
			},
		},
		PostAggregations: &[]intersight.TelemetryDruidPostAggregator{
			{
				TelemetryDruidArithmeticPostAggregator: &intersight.TelemetryDruidArithmeticPostAggregator{
					Type: "arithmetic",
					Name: optStr("bps"),
					Fn:   optStr("/"),
					AdditionalProperties: map[string]interface{}{
						"fields": []interface{}{
							intersight.TelemetryDruidFieldAccessorPostAggregator{
								Type:      "fieldAccess",
								Name:      optStr("traffic"),
								FieldName: optStr("traffic"),
							},
							intersight.TelemetryDruidConstantPostAggregator{
								Type:  "constant",
								Name:  optStr("seconds"),
								Value: optFloat64(112.5),
							},
						},
					},
				},
			},
		},
	}

	c := NewDruidCollector("ucs_ether_port_stat", druidGroupByRequest, []string{"deviceId", "port"})
	c.RegisterMetric("intersight_port_tx_rate", "Port transmit rate (bps)", "bps")
	// c.RegisterMetric("intersight_port_rx_rate", "Port receive rate (bps)", "sumWriteBytes")
	RegisterCollector(c)
}
