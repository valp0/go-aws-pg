package handlers

import (
	"net/http"

	"github.com/valp0/go-aws-pg/services"
)

// PingHandler performs a ping to the database to verify it is available.
func Ping(w http.ResponseWriter, r *http.Request) {
	response, err := services.Ping()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeResponse(w, err)
	}

	w.WriteHeader(http.StatusOK)
	writeResponse(w, response)
}
