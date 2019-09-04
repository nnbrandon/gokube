package handlers

import "net/http"

// liveness probes to help understand that the application is running
// if it fails, the service will be restarted by kubernetes

// healthz is a liveness probe
func healthz(res http.ResponseWriter, _ *http.Request) {
	res.WriteHeader(http.StatusOK)
}
