package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// DeleteUser will delete a user from the users and the user_favs tables, given its id.
func (h handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	users, err := h.s.DeleteUser(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeResponse(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	writeResponse(w, users)
}
