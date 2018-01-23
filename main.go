package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// Definition of a Prometheus counter
	httpCallsCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "http_calls_counter",
			Help: "Counter for HTTP calls",
		},
	)
)

func main() {

	// Adding custom handlers
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		httpCallsCounter.Inc()
		fmt.Fprintf(w, "Counting calls...")
	})

	// Adding prometheus Handler
	http.Handle("/metrics", promhttp.Handler())

	// Starting webserver
	log.Println("Starting WebServer on 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func init() {
	// Every metrics needs to be registered
	prometheus.MustRegister(httpCallsCounter)
}
