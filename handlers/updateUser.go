package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/valp0/go-aws-pg/services"
)

// UpdateUser is the handler to update users via a PATCH request.
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	bodyDecoder := json.NewDecoder(r.Body)

	user, err := services.UpdateUser(id, bodyDecoder)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeResponse(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	writeResponse(w, user)
}
