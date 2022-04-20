package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h handler) DeleteFavorite(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	vidId := mux.Vars(r)["vidId"]

	user, err := h.s.DeleteFavorite(id, vidId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeResponse(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	writeResponse(w, user)
}
