package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/valp0/go-aws-pg/services"
)

// PostUser will respond with the user if it was successfully inserted, or a message if not.
func PostUser(w http.ResponseWriter, r *http.Request) {
	bodyDecoder := json.NewDecoder(r.Body)

	user, err := services.PostUser(bodyDecoder)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeResponse(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	writeResponse(w, user)
}
