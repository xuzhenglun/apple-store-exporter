package metric

import "github.com/prometheus/client_golang/prometheus"

func init() {
	prometheus.MustRegister(
		stockMetricMetric,
	)
}

var (
	stockMetricMetric = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "apple_store",
		Subsystem: "exporter",
		Name:      "fulfillment_messages",
		Help:      "",
	}, []string{"shop", "model"})
)
