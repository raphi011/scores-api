package services

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Metrics struct {
	tournamentSignups *prometheus.CounterVec
}

var m = &Metrics{
	tournamentSignups: promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "api_tournament_signups",
		Help: "The total number of tournament signups",
	}, []string{}),
}

func NewMetrics() *Metrics {
	return m
}
