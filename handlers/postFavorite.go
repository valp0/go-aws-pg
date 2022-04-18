package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/valp0/go-aws-pg/services"
)

// PostFavorite will trigger the functions to add a favorite to favorites and user_favs tables.
func PostFavorite(w http.ResponseWriter, r *http.Request) {
	bodyDecoder := json.NewDecoder(r.Body)
	id := mux.Vars(r)["id"]

	fav, err := services.PostFavorite(id, bodyDecoder)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeResponse(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	writeResponse(w, fav)
}
