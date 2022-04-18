package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/valp0/go-aws-pg/services"
)

func DeleteFavorite(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	vidId := mux.Vars(r)["vidId"]

	user, err := services.DeleteFavorite(id, vidId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeResponse(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	writeResponse(w, user)
}
