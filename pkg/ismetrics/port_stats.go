package ismetrics

import (
	"fmt"
	"time"

	"github.com/cgascoig/intersight-metrics/pkg/intersight"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

func getIntervalString(endTime time.Time, duration time.Duration) string {
	startTime := endTime.Add(-duration)

	layout := "2006-01-02T15:04:05-07:00"

	return fmt.Sprintf("%s/%s", startTime.Format(layout), endTime.Format(layout))
}

func genericCollectPortStats(druidAggReq intersight.TelemetryDruidAggregateRequest, metric string, im *IntersightMetrics, desc *prometheus.Desc, ch chan<- prometheus.Metric) {
	logrus.Infof("Requesting time series data")

	req := im.client.TelemetryApi.QueryTelemetryTimeSeries(im.ctx)

	req = req.TelemetryDruidAggregateRequest(druidAggReq)

	res, _, err := req.Execute()
	if err != nil {
		logrus.Errorf("Error executing timeseries request: %v", err)
	}

	// Go through all the results and make sure we're finding the latest one for each dn/device_id pair
	var results = map[string]intersight.TelemetryDruidIntervalResult{}

	for _, ir := range res {
		var deviceID, dn string
		var event map[string]interface{}
		var ok bool

		if ir.Timestamp == nil {
			logrus.Warn("Skipping interval result without timestamp")
			continue
		}

		if _, ok := ir.AdditionalProperties["event"]; !ok {
			logrus.Warn("Skipping interval result without event")
			continue
		}

		if event, ok = ir.AdditionalProperties["event"].(map[string]interface{}); !ok {
			logrus.Warn("Skipping interval result with invalid event type")
			continue
		}

		if d, ok := event["deviceId"]; !ok {
			logrus.Warn("Skipping interval result without deviceId")
			continue
		} else {
			if deviceID, ok = d.(string); !ok {
				logrus.Warn("Skipping interval result with invalid deviceId")
				continue
			}
		}

		if d, ok := event["dn"]; !ok {
			logrus.Warn("Skipping interval result without dn")
			continue
		} else {
			if dn, ok = d.(string); !ok {
				logrus.Warn("Skipping interval result with invalid dn")
				continue
			}
		}

		if b, ok := event[metric]; !ok {
			logrus.Warn("Skipping interval result without metric ", metric)
			continue
		} else {
			if _, ok = b.(float64); !ok {
				logrus.Warn("Skipping interval result with invalid bytes")
				continue
			}
		}

		k := fmt.Sprintf("%s:%s", deviceID, dn)
		if existing, ok := results[k]; ok {
			if existing.Timestamp.Before(*ir.Timestamp) {
				results[k] = ir
			}
		} else {
			results[k] = ir
		}
	}

	for _, ir := range results {
		event, _ := ir.AdditionalProperties["event"].(map[string]interface{})
		metricVal := event[metric].(float64)
		deviceId := event["deviceId"].(string)
		dn := event["dn"].(string)

		ch <- prometheus.MustNewConstMetric(desc, prometheus.GaugeValue, metricVal, deviceId, dn)

		logrus.Debugf("Updated port: %s %s: %v\n", dn, metric, metricVal)
	}
}

func collectPortBpsStats(im *IntersightMetrics, desc *prometheus.Desc, ch chan<- prometheus.Metric) {
	logrus.Info("Starting port bps stats collection")

	druidAggReq := intersight.TelemetryDruidAggregateRequest{
		TelemetryDruidGroupByRequest: &intersight.TelemetryDruidGroupByRequest{
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
						OutputName: "dn",
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
		},
	}

	genericCollectPortStats(druidAggReq, "bps", im, desc, ch)

}

// func collectPortStats(im *IntersightMetrics, desc *prometheus.Desc, ch chan<- prometheus.Metric) {
// 	logrus.Info("Starting port stats collection")

// 	interval := getIntervalString(time.Now(), time.Duration(15)*time.Minute)
// 	logrus.Infof("Requesting time series data for interval: %s", interval)

// 	req := im.client.TelemetryApi.QueryTelemetryTimeSeries(im.ctx)

// 	req = req.TelemetryDruidAggregateRequest(intersight.TelemetryDruidAggregateRequest{
// 		TelemetryDruidGroupByRequest: &intersight.TelemetryDruidGroupByRequest{
// 			QueryType: "groupBy",
// 			Granularity: intersight.TelemetryDruidGranularity{
// 				TelemetryDruidDurationGranularity: &intersight.TelemetryDruidDurationGranularity{
// 					Type:     "duration",
// 					Duration: 60 * 1000,
// 				},
// 			},
// 			DataSource: intersight.TelemetryDruidDataSource{
// 				TelemetryDruidTableDataSource: &intersight.TelemetryDruidTableDataSource{
// 					Name: "ucs_ether_port_stat",
// 					Type: "table",
// 				},
// 			},
// 			Dimensions: []intersight.TelemetryDruidDimensionSpec{
// 				{
// 					TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
// 						Type:       "default",
// 						Dimension:  "dn",
// 						OutputName: "dn",
// 						OutputType: "STRING",
// 					},
// 				},
// 				{
// 					TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
// 						Type:       "default",
// 						Dimension:  "deviceId",
// 						OutputName: "deviceId",
// 						OutputType: "STRING",
// 					},
// 				},
// 			},
// 			Intervals: []string{interval},
// 			Aggregations: &[]intersight.TelemetryDruidAggregator{
// 				{
// 					TelemetryDruidSumAggregator: &intersight.TelemetryDruidSumAggregator{
// 						Type:      "longSum",
// 						Name:      "traffic",
// 						FieldName: "sumBytesTx",
// 					},
// 				},
// 			},
// 		},
// 	})

// 	res, _, err := req.Execute()
// 	if err != nil {
// 		logrus.Errorf("Error executing timeseries request: %v", err)
// 	}

// 	type MyIntervalResult struct {
// 		DeviceMoid string
// 		Port       string
// 		Bytes      int64
// 		Timestamp  time.Time
// 	}
// 	var results = make(map[string][]MyIntervalResult)

// 	// transform all returned interval results into a map keyed by "deviceID:dn" and skip any with malformed data
// 	for _, ir := range res {
// 		var deviceID string
// 		var dn string
// 		var timestamp time.Time
// 		var bytes float64
// 		var event map[string]interface{}
// 		var ok bool

// 		if ir.Timestamp == nil {
// 			logrus.Warn("Skipping interval result without timestamp")
// 			continue
// 		}
// 		timestamp = *ir.Timestamp

// 		if _, ok := ir.AdditionalProperties["event"]; !ok {
// 			logrus.Warn("Skipping interval result without event")
// 			continue
// 		}

// 		if event, ok = ir.AdditionalProperties["event"].(map[string]interface{}); !ok {
// 			logrus.Warn("Skipping interval result with invalid event type")
// 			continue
// 		}

// 		if d, ok := event["deviceId"]; !ok {
// 			logrus.Warn("Skipping interval result without deviceId")
// 			continue
// 		} else {
// 			if deviceID, ok = d.(string); !ok {
// 				logrus.Warn("Skipping interval result with invalid deviceId")
// 				continue
// 			}
// 		}

// 		if d, ok := event["dn"]; !ok {
// 			logrus.Warn("Skipping interval result without dn")
// 			continue
// 		} else {
// 			if dn, ok = d.(string); !ok {
// 				logrus.Warn("Skipping interval result with invalid dn")
// 				continue
// 			}
// 		}

// 		if b, ok := event["traffic"]; !ok {
// 			logrus.Warn("Skipping interval result without traffic")
// 			continue
// 		} else {
// 			if bytes, ok = b.(float64); !ok {
// 				logrus.Warn("Skipping interval result with invalid bytes")
// 				continue
// 			}
// 		}

// 		logrus.Debugf("Interval result: time=%v deviceId=%s dn=%s bytes=%v", timestamp, deviceID, dn, int64(bytes))
// 		k := fmt.Sprintf("%s:%s", deviceID, dn)
// 		results[k] = append(results[k], MyIntervalResult{
// 			DeviceMoid: deviceID,
// 			Port:       dn,
// 			Timestamp:  timestamp,
// 			Bytes:      int64(bytes),
// 		})
// 	}

// 	// go through all the transformed results and calculate stats per port(deviceID:dn)
// 	for port, irs := range results {
// 		sort.Slice(irs, func(i, j int) bool {
// 			return irs[i].Timestamp.Before(irs[j].Timestamp)
// 		})

// 		// for _, ir := range irs {
// 		// 	fmt.Printf(" %v: %v\n", ir.Timestamp, ir.Bytes)
// 		// }

// 		numIr := len(irs)
// 		if numIr < 2 {
// 			logrus.Warnf("Skipping %s due to insufficient data", port)
// 			continue
// 		}
// 		td := irs[numIr-1].Timestamp.Sub(irs[0].Timestamp)

// 		var totalBytes float64 = 0
// 		for i := 0; i < numIr; i++ {
// 			totalBytes += float64(irs[i].Bytes)
// 		}

// 		bps := totalBytes * 8 / td.Seconds()

// 		ch <- prometheus.MustNewConstMetric(desc, prometheus.GaugeValue, bps, irs[numIr-2].DeviceMoid, irs[numIr-2].Port)

// 		logrus.Debugf("Updated port: %s bps: %v\n", port, bps)
// 	}

// 	logrus.Info("Finished port stats collection")
// }

func init() {
	RegisterMetric("intersight_port_tx_rate", "Port transmit rate (bps)", []string{"device_moid", "port"}, collectPortBpsStats)
}
