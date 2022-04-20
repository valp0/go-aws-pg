package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// UpdateUser is the handler to update users via a PATCH request.
func (h handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	bodyDecoder := json.NewDecoder(r.Body)

	user, err := h.s.UpdateUser(id, bodyDecoder)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeResponse(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	writeResponse(w, user)
}
