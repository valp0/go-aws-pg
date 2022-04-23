package handlers

import (
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

// GetUsers will fetch all users in the users table.
func (h handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	token := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)

	claims := token.CustomClaims.(*CustomClaims)
	if !claims.HasScope("read:users") {
		w.WriteHeader(http.StatusForbidden)
		writeResponse(w, "Insufficient scope.")
		return
	}

	users, err := h.s.GetUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeResponse(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	writeResponse(w, users)
}
