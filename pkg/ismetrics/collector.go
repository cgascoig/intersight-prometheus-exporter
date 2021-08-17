package ismetrics

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
)

type IntersightCollector interface {
	Collect(im *IntersightMetrics, ctx context.Context, ch chan<- prometheus.Metric)
	Describe(ch chan<- *prometheus.Desc)
}
