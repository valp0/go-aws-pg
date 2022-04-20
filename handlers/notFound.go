package handlers

import (
	"net/http"
)

// NotFound will send a message inside a JSON containing the available endpoints.
func (h handler) NotFound(w http.ResponseWriter, r *http.Request) {
	message := `only available endpoints are GET /api/ping/, GET/POST/DELETE /api/users/, GET/PATCH /api/users/{id}, GET/POST /api/users/{id}/favorites/ and DELETE /api/users/{id}/favorites{vidId}/`

	w.WriteHeader(http.StatusNotFound)
	writeResponse(w, message)
}
