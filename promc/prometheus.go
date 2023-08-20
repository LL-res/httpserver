package promc

import (
	"github.com/prometheus/client_golang/prometheus"
	"httpserver/consts"
)

var (
	QPSCounter = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "http_requests_total",
		Help: "Number of HTTP requests processed.",
	}, []string{consts.METHOD, consts.STATUS_CODE})
	T3Counter = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "http_requests_nums",
		Help: "Number of HTTP requests processed for every 30 seconds.",
	}, []string{consts.METHOD, consts.STATUS_CODE})
	T3Flag = true
)

func Counters() []prometheus.Collector {
	counters := make([]prometheus.Collector, 0)

	counters = append(counters, QPSCounter, T3Counter)

	return counters
}
func Init() {
	prometheus.MustRegister(Counters()...)
}
