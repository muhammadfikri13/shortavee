package metrics

import "github.com/prometheus/client_golang/prometheus"

var URLsCreated = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "shortavee_urls_created_total",
		Help: "Total URLs created",
	},
)

var Redirects = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "shortavee_redirects_total",
		Help: "Total redirects",
	},
)

func Register() {
	prometheus.MustRegister(URLsCreated)
	prometheus.MustRegister(Redirects)
}
