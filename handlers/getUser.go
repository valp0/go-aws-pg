package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// GetUser fetches a user by its id.
func (h handler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	user, err := h.s.GetUser(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeResponse(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	writeResponse(w, user)
}
