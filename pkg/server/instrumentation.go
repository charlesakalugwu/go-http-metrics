package server

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	authenticatedUserCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "authenticated_user_requests",
			Help: "Counter of authenticated requests broken out by username.",
		},
		[]string{"username"},
	)

	helloWorldsCounter = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "hello_worlds_total",
			Help: "Hello worlds requested.",
		},
	)

	inProgressGauge = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "hello_worlds_inprogress",
			Help: "Number of hello worlds in progress.",
		},
	)

	lastCallGauge = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "hello_worlds_last_time_seconds",
			Help: "The last time a hello world was served.",
		},
	)

	latencySummary = promauto.NewSummary(
		prometheus.SummaryOpts{
			Name: "hello_world_latency_seconds",
			Help: "Time to request a hello world.",
		},
	)

	latencyHistogram = promauto.NewHistogram(
		prometheus.HistogramOpts{
			Name: "app_latency_seconds",
			Help: "Time to request a hello world.",
		},
	)
)
