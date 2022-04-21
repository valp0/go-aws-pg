package handlers

import (
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gorilla/mux"
)

// GetFavorites fetches the favorites of a user by its id.
func (h handler) GetFavorites(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	w.Header().Set("Content-Type", "application/json")

	token := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)

	claims := token.CustomClaims.(*CustomClaims)
	if !claims.HasScope("read:favorites") {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(`{"message":"Insufficient scope."}`))
		return
	}

	favs, err := h.s.GetFavorites(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeResponse(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	writeResponse(w, favs)
}
