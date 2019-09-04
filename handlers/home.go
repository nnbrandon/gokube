package handlers

import (
	"encoding/json"
	"github.com/mrnguuyen/go_kube/version"
	"log"
	"net/http"
)

// home is a simple HTTP handler function which writes a response
func home(buildTime, commit, release string) http.HandlerFunc {
	return func(res http.ResponseWriter, _ *http.Request) {
		info := struct {
			BuildTime string `json:"buildTime"`
			Commit	  string `json:"commit"`
			Release   string `json:"release"`
		} {
			version.BuildTime, version.Commit, version.Release,
		}

		body, err := json.Marshal(info)
		if err != nil {
			log.Printf("Could not encode info data %v", err)
			http.Error(res, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
			return
		}
		res.Header().Set("Content-Type", "application/json")
		res.Write(body)
	}
}