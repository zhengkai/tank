package metrics

import "github.com/prometheus/client_golang/prometheus"

func init() {
	prometheus.MustRegister(crawlTime)
	prometheus.MustRegister(crawlFail)
}
