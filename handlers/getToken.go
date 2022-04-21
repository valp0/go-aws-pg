package handlers

import (
	"encoding/json"
	"net/http"
)

// GetToken will respond with the currently valid token and its remaining time.
func (h handler) GetToken(w http.ResponseWriter, r *http.Request) {
	bodyDecoder := json.NewDecoder(r.Body)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080/")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization")

	token, err := h.s.GetToken(bodyDecoder)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeResponse(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	writeResponse(w, &token)
}
