package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Receives a http.ResponseWriter and an result, then writes a response depending on the result type.
func writeResponse(w http.ResponseWriter, result interface{}) {
	var response Data
	if _, isStr := result.(string); isStr {
		response = Data{Message{result}}
	} else {
		response = Data{Items{result}}
	}

	jResponse, _ := json.Marshal(response)
	fmt.Fprintln(w, prettifyJson(jResponse))
}

// Prettifies a JSON byte array and returns a prettified JSON string.
func prettifyJson(ugly []byte) string {
	var prettified bytes.Buffer
	err := json.Indent(&prettified, ugly, "", "  ")
	if err != nil {
		return string(ugly)
	}

	return prettified.String()
}
