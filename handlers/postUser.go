package handlers

import (
	"encoding/json"
	"net/http"
)

// PostUser will respond with the user if it was successfully inserted, or a message if not.
func (h handler) PostUser(w http.ResponseWriter, r *http.Request) {
	bodyDecoder := json.NewDecoder(r.Body)

	user, err := h.s.PostUser(bodyDecoder)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeResponse(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	writeResponse(w, user)
}
