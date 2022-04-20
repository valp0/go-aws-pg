package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// GetFavorites fetches the favorites of a user by its id.
func (h handler) GetFavorites(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	favs, err := h.s.GetFavorites(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeResponse(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	writeResponse(w, favs)
}
