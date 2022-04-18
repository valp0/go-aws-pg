package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/valp0/go-aws-pg/services"
)

// GetUser fetches a user by its id.
func GetUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	user, err := services.GetUser(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeResponse(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	writeResponse(w, user)
}
