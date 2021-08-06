package ismetrics

import "github.com/prometheus/client_golang/prometheus"

type IntersightCollector interface {
	Collect(im *IntersightMetrics, ch chan<- prometheus.Metric)
	Describe(ch chan<- *prometheus.Desc)
}
