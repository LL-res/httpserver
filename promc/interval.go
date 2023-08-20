package promc

import (
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

func ResetByInterval(interval time.Duration, gauge *prometheus.GaugeVec, first *bool) {
	if !*first {
		return
	}

	go func(interval time.Duration, gauge *prometheus.GaugeVec) {
		ticker := time.NewTicker(interval)
		for _ = range ticker.C {
			gauge.Reset()
		}
	}(interval, gauge)

	*first = false
	return
}
