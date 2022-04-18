package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/valp0/go-aws-pg/services"
)

// GetFavorites fetches the favorites of a user by its id.
func GetFavorites(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	favs, err := services.GetFavorites(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeResponse(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	writeResponse(w, favs)
}
