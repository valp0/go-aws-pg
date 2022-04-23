package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/valp0/go-aws-pg/services"
)

// Receives a http.ResponseWriter and an result, then writes a response depending on the result type.
func writeResponse(w http.ResponseWriter, result interface{}) {
	var response Data
	if _, isStr := result.(string); isStr {
		response = Data{Message{result}}
	} else if _, isToken := result.(services.Client); isToken {
		response = Data{Token{result}}
	} else {
		response = Data{Items{result}}
	}

	jResponse, _ := json.Marshal(response)
	w.Write(jResponse)
}
