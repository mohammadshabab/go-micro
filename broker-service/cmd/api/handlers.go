package main

import (
	"net/http"
)

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := response{
		Error:   false,
		Message: "I am broker",
	}
	_ = app.writeJSON(w, http.StatusOK, payload)
	// out, _ := json.MarshalIndent(payload, "", "\t")
	// slog.Debug("Writing paylod ", "payload ", payload)
	// w.Header().Add("Content-Type", "application/json")
	// w.WriteHeader(http.StatusAccepted)
	// w.Write(out)
}
