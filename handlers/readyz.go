package handlers

import (
	"net/http"
	"sync/atomic"
)

// readiness probe: sometimes, we need to wait for some event to be able to serve traffic

// readyz is a readiness probe
func readyz(isReady *atomic.Value) http.HandlerFunc {
	return func(res http.ResponseWriter, _ *http.Request) {
		if isReady == nil || isReady.Load().(bool) {
			http.Error(res, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
			return
		}
		// return 200 only if isReady is set and equals to true
		res.WriteHeader(http.StatusOK)
	}
}