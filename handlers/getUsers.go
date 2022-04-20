package handlers

import (
	"net/http"
)

// GetUsers fetches all users in the users table.
func (h handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.s.GetUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeResponse(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	writeResponse(w, users)
}
