package ismetrics

import (
	"context"
	"time"

	"github.com/cgascoig/intersight-prometheus-exporter/pkg/intersight"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

type DruidMetric struct {
	desc        *prometheus.Desc
	druidMetric string
}

type DruidCollector struct {
	name string

	metrics    []*DruidMetric
	labelNames []string

	druidGroupByRequest intersight.TelemetryDruidGroupByRequest
}

func NewDruidCollector(name string, druidGroupByRequest intersight.TelemetryDruidGroupByRequest, labelNames []string) *DruidCollector {
	return &DruidCollector{
		name:                name,
		labelNames:          labelNames,
		druidGroupByRequest: druidGroupByRequest,
	}
}

func (c *DruidCollector) RegisterMetric(prometheusMetricName, prometheusMetricHelp, druidMetric string) {
	dm := &DruidMetric{
		desc:        prometheus.NewDesc(prometheusMetricName, prometheusMetricHelp, c.labelNames, nil),
		druidMetric: druidMetric,
	}

	c.metrics = append(c.metrics, dm)
}

func (c *DruidCollector) Describe(ch chan<- *prometheus.Desc) {
	for _, dm := range c.metrics {
		ch <- dm.desc
	}
}

func (c *DruidCollector) Collect(im *IntersightMetrics, ctx context.Context, ch chan<- prometheus.Metric) {
	logrus.Info("Requesting time series data for ", c.name)

	c.druidGroupByRequest.Intervals = []string{getIntervalString(time.Now(), time.Duration(15)*time.Minute)}

	req := im.client.TelemetryApi.QueryTelemetryGroupBy(ctx)

	req = req.TelemetryDruidGroupByRequest(c.druidGroupByRequest)

	res, _, err := req.Execute()
	if err != nil {
		logrus.Errorf("Error executing timeseries request: %v", err)
	}

	// Go through all the results and make sure we're finding the latest one for each dn/device_id pair
	var results = map[string]intersight.TelemetryDruidGroupByResult{}

checkIRs:
	for _, ir := range res {
		if ir.Timestamp == nil {
			logrus.Warn("Skipping interval result without timestamp")
			continue
		}

		// Check that the IR has the required dimensions/labelNames
		key := ""
		event := *ir.Event
		for _, labelName := range c.labelNames {
			if d, ok := event[labelName]; !ok {
				logrus.Warn("Skipping interval result without dimension ", labelName)
				continue checkIRs
			} else {
				if label, ok := d.(string); !ok {
					logrus.Warn("Skipping interval result with invalid dimension ", labelName)
					continue checkIRs
				} else {
					key += label + ":"
				}
			}
		}

		if existing, ok := results[key]; ok {
			if existing.Timestamp.Before(*ir.Timestamp) {
				results[key] = ir
			}
		} else {
			results[key] = ir
		}
	}

rangeResults:
	for _, ir := range results {
		event := *ir.Event
		for _, metric := range c.metrics {
			var metricVal float64
			var labels []string

			if b, ok := event[metric.druidMetric]; !ok {
				continue rangeResults
			} else {
				if metricVal, ok = b.(float64); !ok {
					continue rangeResults
				}
			}

			for _, labelName := range c.labelNames {
				if d, ok := event[labelName]; !ok {
					continue rangeResults
				} else {
					if label, ok := d.(string); !ok {
						continue rangeResults
					} else {
						labels = append(labels, label)
					}
				}
			}

			ch <- prometheus.MustNewConstMetric(metric.desc, prometheus.GaugeValue, metricVal, labels...)

			logrus.Debugf("Updated metric: %s %s{%v}: %v\n", metric.desc.String(), metric.druidMetric, labels, metricVal)
		}
	}
}
