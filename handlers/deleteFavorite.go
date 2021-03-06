package handlers

import (
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gorilla/mux"
)

// DeleteFavorite deletes a favorite from the user_favs table, then checks if it is still
// linked to another user after deletion and deletes it from the favorites table if not.
func (h handler) DeleteFavorite(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	vidId := mux.Vars(r)["vidId"]
	w.Header().Set("Content-Type", "application/json")

	token := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)

	claims := token.CustomClaims.(*CustomClaims)
	if !claims.HasScope("read:favorites") || !claims.HasScope("write:favorites") {
		w.WriteHeader(http.StatusForbidden)
		writeResponse(w, "Insufficient scope.")
		return
	}

	user, err := h.s.DeleteFavorite(id, vidId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeResponse(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	writeResponse(w, user)
}
