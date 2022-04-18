package handlers

import (
	"net/http"

	"github.com/valp0/go-aws-pg/services"
)

// GetUsers fetches all users in the users table.
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := services.GetUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeResponse(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	writeResponse(w, users)
}
