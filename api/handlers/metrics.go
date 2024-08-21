package handlers

import (
	"expvar"
	"net/http"
)

var (
	RequestCount  = expvar.NewInt("RequestCount")
	SuccessCount  = expvar.NewInt("SuccessCount")
	ErrorCount    = expvar.NewInt("ErrorCount")
)

func MetricsHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		RequestCount.Add(1)
		if r.URL.Path == "/metrics" {
			expvar.Handler().ServeHTTP(w, r)
			SuccessCount.Add(1)
		} else {
			ErrorCount.Add(1)
			http.NotFound(w, r)
		}
	})
}
