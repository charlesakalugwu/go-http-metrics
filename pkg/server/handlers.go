package server

import (
	"net/http"
	"time"
)

func handleGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
}

func handleHelloWorld(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	helloWorldsCounter.Inc()
	inProgressGauge.Inc()
	w.Write([]byte("Hello world"))
	lastCallGauge.SetToCurrentTime()
	inProgressGauge.Dec()
	latencySummary.Observe(time.Now().Sub(start).Seconds())
	latencyHistogram.Observe(time.Now().Sub(start).Seconds())
}
