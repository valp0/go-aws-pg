package handlers

import (
	"net/http"
)

// PingHandler will perform a ping to the database to verify it is available.
func (h handler) Ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response, err := h.s.Ping()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeResponse(w, err)
	}

	w.WriteHeader(http.StatusOK)
	writeResponse(w, response)
}
